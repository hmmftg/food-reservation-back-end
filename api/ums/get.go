package ums

import (
	"log"
	"net/http"

	"github.com/hmmftg/food-reservation-back-end/internal/otp"
	"github.com/hmmftg/requestCore/handlers"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/libRequest"
)

func (u User) GetCheckData() *CheckResponse {
	return &CheckResponse{
		UserId:        u.ID,
		UserName:      u.Name,
		PersonID:      u.PersonID,
		Roles:         GetRoles(u.Roles),
		Authenticated: true,
	}
}

func (u User) GetPermissonData() *CheckResponse {
	return &CheckResponse{
		Roles: GetRoles(u.Roles),
	}
}

func (u User) GetIDData() *CheckResponse {
	return &CheckResponse{
		UserId:        u.ID,
		UserName:      u.Name,
		PersonID:      u.PersonID,
		Authenticated: true,
		Roles:         GetRoles(u.Roles),
	}
}
func (env umsEnv) umsCheck(simulation bool) any {
	otp, err := otp.GetOTP(env.Interface.Params())
	if err != nil {
		log.Fatalln(err)
	}
	return handlers.BaseHandler(env.Interface, CheckHandler{Name: "ums-check", HMac: otp}, simulation)
}

func (env umsEnv) umsPermissions(simulation bool) any {
	otp, err := otp.GetOTP(env.Interface.Params())
	if err != nil {
		log.Fatalln(err)
	}
	return handlers.BaseHandler(env.Interface, CheckHandler{Name: "ums-perm", HMac: otp}, simulation)
}

func (env umsEnv) umsGetUser(simulation bool) any {
	otp, err := otp.GetOTP(env.Interface.Params())
	if err != nil {
		log.Fatalln(err)
	}
	return handlers.BaseHandler(env.Interface, CheckHandler{Name: "ums-get-id", HMac: otp}, simulation)
}

// returns handler title
//
//	Request Bodymode
//	and validate header option
//	and save to request table option
//	and url path of handler
func (h CheckHandler) Parameters() handlers.HandlerParameters {
	return handlers.HandlerParameters{
		Title:          "ums",
		Body:           libRequest.NoBinding,
		ValidateHeader: true,
		SaveToRequest:  false,
		Path:           "/ums",
	}
}

// runs after validating request
func (h CheckHandler) Initializer(req handlers.HandlerRequest[CheckRequest, *CheckResponse]) error {
	return nil
}

// Handler is the main method that handles request and returns the response,
// if there is a need for calling another api this is the place to call that api.
func (h CheckHandler) Handler(req handlers.HandlerRequest[CheckRequest, *CheckResponse]) (*CheckResponse, error) {
	token, err := GetToken(req.W)
	if err != nil {
		return nil, err
	}
	usr, err := ValidateJwtToken(req.Core, token)
	if err != nil {
		return nil, err
	}
	err = otp.Check(h.HMac, usr.ID, req.Core, req.W)
	if err != nil {
		return nil, err
	}

	switch h.Name {
	case "ums-check":
		return usr.GetCheckData(), nil
	case "ums-perm":
		return usr.GetPermissonData(), nil
	case "ums-get-id":
		return usr.GetIDData(), nil
	}
	return nil, libError.NewWithDescription(http.StatusInternalServerError, "UNKNOWN_METHOD", "method not defined: %s", h.Name)
}

// Simulation returns a simulated response.
func (h CheckHandler) Simulation(req handlers.HandlerRequest[CheckRequest, *CheckResponse]) (*CheckResponse, error) {
	return req.Response, nil
}

// runs after sending back response
func (h CheckHandler) Finalizer(req handlers.HandlerRequest[CheckRequest, *CheckResponse]) {
}
