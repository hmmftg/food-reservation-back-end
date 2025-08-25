package ums

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hmmftg/food-reservation-back-end/internal/otp"
	"github.com/hmmftg/food-reservation-back-end/internal/password"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/handlers"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/libQuery"
	"github.com/hmmftg/requestCore/libRequest"
	"github.com/hmmftg/requestCore/webFramework"
)

func (u UserData) GetLoginData(token, refreshToken string) *LoginResponse {
	return &LoginResponse{
		ID:           u.ID,
		Name:         u.Name,
		Department:   u.Department,
		Roles:        u.Roles,
		PersonID:     u.PersonID,
		UserData:     u.UserData,
		RefreshToken: refreshToken,
		AccessToken:  token,
	}
}

func (env umsEnv) umsLogin(simulation bool) any {
	otp, err := otp.GetOTP(env.Interface.Params())
	if err != nil {
		log.Fatalln(err)
	}
	return handlers.BaseHandler(env.Interface, LoginHandler{Name: "ums-login", HMac: otp}, simulation)
}

// returns handler title
//
//	Request Bodymode
//	and validate header option
//	and save to request table option
//	and url path of handler
func (h LoginHandler) Parameters() handlers.HandlerParameters {
	return handlers.HandlerParameters{
		Title:          "ums",
		Body:           libRequest.NoBinding,
		ValidateHeader: false,
		SaveToRequest:  false,
		Path:           "/ums",
	}
}

// runs after validating request
func (h LoginHandler) Initializer(req handlers.HandlerRequest[LoginRequest, *LoginResponse]) error {
	return nil
}

func GetUserData(userName string, core requestCore.RequestCoreInterface) (*UserData, error) {
	result, err := libQuery.GetQuery[UserData](`--sql
		select * 
		  from [USERS] 
		 where id = ?
	`, core.GetDB(), userName)
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func GetUserPass(w webFramework.WebFramework) (*LoginRequest, error) {
	authHeader := w.Parser.GetHeaderValue("Authorization")
	if len(authHeader) == 0 {
		return nil, libError.NewWithDescription(http.StatusBadRequest,
			"AUTH_HEADER_ABSENT_OR_INVALID", "auth header does not exists")
	}
	auth := strings.SplitN(authHeader, " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return nil, libError.NewWithDescription(http.StatusUnauthorized, "AUTH_BAD_METHOD", "auth header type is not valid")
	}
	userPassByte, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return nil, libError.New(http.StatusBadRequest,
			"AUTH_HEADER_INVALID_FORMAT", err.Error())
	}
	userPass := strings.Split(string(userPassByte), ":")
	return &LoginRequest{UserName: userPass[0], Pass: userPass[1]}, nil

}

// Handler is the main method that handles request and returns the response,
// if there is a need for calling another api this is the place to call that api.
func (h LoginHandler) Handler(req handlers.HandlerRequest[LoginRequest, *LoginResponse]) (*LoginResponse, error) {
	switch h.Name {
	case "ums-login":
		var errH error
		req.Request, errH = GetUserPass(req.W)
		if errH != nil {
			return nil, errH
		}
		errLocker := otp.Check(h.HMac, req.Request.UserName, req.Core, req.W)
		if errLocker != nil {
			return nil, errLocker
		}
		user, err := GetUserData(req.Request.UserName, req.Core)
		if err != nil {
			return nil, libError.New(http.StatusBadRequest, "USER_NOT_FOUND", err.Error())
		}
		if password.GetHash3(req.Request.Pass) != user.Password {
			return nil, libError.NewWithDescription(http.StatusBadRequest, "INVALID_PASSWORD",
				"invalid password")
		}

		dt := time.Now().UTC()

		token, err := GenerateToken(dt,
			36000,
			"simple",
			user.Roles,
			user.ID)
		if err != nil {
			return nil, libError.New(http.StatusBadRequest, "ERROR_GENERATE_TOKEN", err.Error())
		}

		refreshToken, err := GenerateToken(dt,
			72000,
			"simple",
			user.Roles,
			user.ID)
		if err != nil {
			return nil, libError.New(http.StatusBadRequest, "ERROR_GENERATE_TOKEN", err.Error())
		}
		return user.GetLoginData(token, refreshToken), nil
	}
	return nil, libError.NewWithDescription(
		http.StatusInternalServerError,
		"UNKNOWN_METHOD",
		"method not defined: %s", h.Name)
}

// Simulation returns a simulated response.
func (h LoginHandler) Simulation(req handlers.HandlerRequest[LoginRequest, *LoginResponse]) (*LoginResponse, error) {
	return req.Response, nil
}

// runs after sending back response
func (h LoginHandler) Finalizer(req handlers.HandlerRequest[LoginRequest, *LoginResponse]) {
}
