package ums

import (
	"net/http"

	"github.com/gofiber/fiber/v2/log"
	"github.com/hmmftg/requestCore/handlers"
	"github.com/hmmftg/requestCore/libError"
	"github.com/hmmftg/requestCore/libRequest"
)

func (env umsEnv) umsLogout(simulation bool) any {
	return handlers.BaseHandler(env.Interface, LogoutHandler{Name: "ums-logout"}, simulation)
}

// returns handler title
//
//	Request Bodymode
//	and validate header option
//	and save to request table option
//	and url path of handler
func (h LogoutHandler) Parameters() handlers.HandlerParameters {
	return handlers.HandlerParameters{
		Title:          "ums",
		Body:           libRequest.JSON,
		ValidateHeader: true,
		SaveToRequest:  false,
		Path:           "/ums",
	}
}

// runs after validating request
func (h LogoutHandler) Initializer(req handlers.HandlerRequest[LogoutRequest, *LogoutResponse]) error {
	return nil
}

// Handler is the main method that handles request and returns the response,
// if there is a need for calling another api this is the place to call that api.
func (h LogoutHandler) Handler(req handlers.HandlerRequest[LogoutRequest, *LogoutResponse]) (*LogoutResponse, error) {
	switch h.Name {
	case "ums-logout":
		token, err := GetToken(req.W)
		if err != nil {
			log.Error(err)
			return &LogoutResponse{State: "error-get-token"}, nil
		}
		_, err = ValidateJwtToken(req.Core, token)
		if err != nil {
			log.Error(err)
			return &LogoutResponse{State: "error-validate-token"}, nil
		}

		return &LogoutResponse{State: "ok"}, nil
	}
	return nil, libError.NewWithDescription(http.StatusInternalServerError, "UNKNOWN_METHOD", "method not defined: %s", h.Name)
}

// Simulation returns a simulated response.
func (h LogoutHandler) Simulation(req handlers.HandlerRequest[LogoutRequest, *LogoutResponse]) (*LogoutResponse, error) {
	return req.Response, nil
}

// runs after sending back response
func (h LogoutHandler) Finalizer(req handlers.HandlerRequest[LogoutRequest, *LogoutResponse]) {
}
