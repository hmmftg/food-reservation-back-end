package ums

import (
	"database/sql"
	"net/http"

	"github.com/hmmftg/food-reservation-back-end/internal/password"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/handlers"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/libRequest"
)

func (env umsEnv) umsRegister(simulation bool) any {
	return handlers.BaseHandler(env.Interface, RegisterHandler{Name: "ums-register"}, simulation)
}

// returns handler title
//
//	Request Bodymode
//	and validate header option
//	and save to request table option
//	and url path of handler
func (h RegisterHandler) Parameters() handlers.HandlerParameters {
	return handlers.HandlerParameters{
		Title:          "ums",
		Body:           libRequest.JSON,
		ValidateHeader: false,
		SaveToRequest:  false,
		Path:           "/ums",
	}
}

// runs after validating request
func (h RegisterHandler) Initializer(req handlers.HandlerRequest[RegisterRequest, *RegisterResponse]) error {
	return nil
}

func InsertUserData(userID, userName, pass string, core requestCore.RequestCoreInterface) (sql.Result, error) {
	result, err := core.GetDB().InsertRow(`--sql
		insert into [USERS] values(?, ?, NULL, '000014', NULL, NULL, ?)
	`, userID, userName, pass)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Handler is the main method that handles request and returns the response,
// if there is a need for calling another api this is the place to call that api.
func (h RegisterHandler) Handler(req handlers.HandlerRequest[RegisterRequest, *RegisterResponse]) (*RegisterResponse, error) {
	switch h.Name {
	case "ums-register":
		_, err := GetUserData(req.Request.UserID, req.Core)
		if err == nil {
			return nil, libError.New(http.StatusBadRequest, "USER_EXISTS", req.Request.UserID)
		}
		req.Request.Pass = password.GetHash3(req.Request.Pass)

		_, err = InsertUserData(req.Request.UserID, req.Request.UserName, req.Request.Pass, req.Core)
		if err != nil {
			return nil, libError.New(http.StatusBadRequest, "ERROR_INSERT_NEW_USER", err.Error())
		}
		return &RegisterResponse{}, nil

	}
	return nil, libError.NewWithDescription(
		http.StatusInternalServerError,
		"UNKNOWN_METHOD",
		"method not defined: %s", h.Name)
}

// Simulation returns a simulated response.
func (h RegisterHandler) Simulation(req handlers.HandlerRequest[RegisterRequest, *RegisterResponse]) (*RegisterResponse, error) {
	return req.Response, nil
}

// runs after sending back response
func (h RegisterHandler) Finalizer(req handlers.HandlerRequest[RegisterRequest, *RegisterResponse]) {
}
