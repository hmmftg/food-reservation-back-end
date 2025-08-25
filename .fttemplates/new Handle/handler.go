package [FTName|camelcase?lowercase]

import (
	"fmt"
	"net/http"

	"github.com/hmmftg/requestCore/handlers"
	"github.com/hmmftg/requestCore/libQuery"
	"github.com/hmmftg/requestCore/libRequest"
	"github.com/hmmftg/requestCore/response"
)

// returns handler title
//    
//   Request Bodymode
//   and validate header option
//   and save to request table option
//   and url path of handler
func (h [FTName|pascalcase]Handler) Parameters() handlers.HandlerParameters {
	return handlers.HandlerParameters{
		Title:          "[FTName|capitalcase]",
		Body:           libRequest.JSON,
		ValidateHeader: true,
		SaveToRequest:  false,
		Path:           "/[FTName|paramcase]",
	}
}

// runs after validating request
func (h [FTName|pascalcase]Handler) Initializer(req handlers.HandlerRequest[[FTName|pascalcase]Request, *[FTName|pascalcase]Response]) error {
	return nil
}

// Handler is the main method that handles request and returns the response,
// if there is a need for calling another api this is the place to call that api.
func (h [FTName|pascalcase]Handler) Handler(req handlers.HandlerRequest[[FTName|pascalcase]Request, *[FTName|pascalcase]Response]) (*[FTName|pascalcase]Response, error) {
	switch h.Name {
	case "[FTName|kebabcase]-post":
		result, err := req.Core.GetDB().InsertRow(`--sql
			insert into SIMULATOR.[FTName|constantcase] (id, name)
			values(:1, :2)
		`, req.Request.ID, req.Request.Name)
		if err != nil {
			return nil, libError.New(http.StatusInternalServerError, "ERROR_INSERT", err.Error())
		}
		dmlResult := libQuery.GetDmlResult(result, nil)
		req.Response = &[FTName|pascalcase]Response{
			Result: dmlResult,
		}
		return req.Response, nil
	case "[FTName|kebabcase]-put":
		result, err := req.Core.GetDB().InsertRow(`--sql
			update SIMULATOR.[FTName|constantcase]
			   set name=:1
			 where id=:2
		`, req.Request.Name, req.Request.ID)
		if err != nil {
			return nil, libError.New(http.StatusInternalServerError, "ERROR_UPDATE", err.Error())
		}
		dmlResult := libQuery.GetDmlResult(result, nil)
		req.Response = &[FTName|pascalcase]Response{
			Result: dmlResult,
		}
		return req.Response, nil
	case "[FTName|kebabcase]-delete":
		result, err := req.Core.GetDB().InsertRow(`--sql
			delete from SIMULATOR.[FTName|constantcase]
			 where id=:1
		`, req.Request.ID)
		if err != nil {
			return nil, libError.New(http.StatusInternalServerError, "ERROR_DELETE", err.Error())
		}
		dmlResult := libQuery.GetDmlResult(result, nil)
		req.Response = &[FTName|pascalcase]Response{
			Result: dmlResult,
		}
		return req.Response, nil
	}
	return nil, libError.NewWithDescription(http.StatusInternalServerError, "UNKNOWN_METHOD", "method not defined: %s", h.Name)
}

// Simulation returns a simulated response.
func (h [FTName|pascalcase]Handler) Simulation(req handlers.HandlerRequest[[FTName|pascalcase]Request, *[FTName|pascalcase]Response]) (*[FTName|pascalcase]Response, error) {
	return req.Response, nil
}

// runs after sending back response
func (h [FTName|pascalcase]Handler) Finalizer(req handlers.HandlerRequest[[FTName|pascalcase]Request, *[FTName|pascalcase]Response]) {
}

// [FTName]PostHandler godoc
// @Summary [FTName]PostHandler
// @Schemes
// @Description  [FTName]PostHandler
// @Tags [FTName]
// @Accept json
// @Produce json
// @Param Request-Id header string true "شناسه درخواست"
// @Param ورودی body [FTName|pascalcase]Request true "اطلاعات ورودی"
// @Router /[FTName|kebabcase] [post]
// @Security  OAuth2Password
// @Success 200 {object} [FTName|pascalcase]Response
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
func (env [FTName|camelcase]Env) [FTName|pascalcase]PostHandler(simulation bool) any {
	return handlers.BaseHandler[[FTName|pascalcase]Request, *[FTName|pascalcase]Response, [FTName|pascalcase]Handler](
		env.Interface, 
		[FTName|pascalcase]Handler{Name:"[FTName|kebabcase]-post"}, 
		simulation,
	)
}


// [FTName]PutHandler godoc
// @Summary [FTName]PutHandler
// @Schemes
// @Description  [FTName]PutHandler
// @Tags [FTName]
// @Accept json
// @Produce json
// @Param Request-Id header string true "شناسه درخواست"
// @Param ورودی body [FTName|pascalcase]Request true "اطلاعات ورودی"
// @Param id path string true "شناسه"
// @Router /[FTName|kebabcase]/:id [put]
// @Security  OAuth2Password
// @Success 200 {object} [FTName|pascalcase]Response
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
func (env [FTName|camelcase]Env) [FTName|pascalcase]PutHandler(simulation bool) any {
	return handlers.BaseHandler[[FTName|pascalcase]Request, *[FTName|pascalcase]Response, [FTName|pascalcase]Handler](
		env.Interface, 
		[FTName|pascalcase]Handler{Name:"[FTName|kebabcase]-put"}, 
		simulation,
	)
}


// [FTName]PutHandler godoc
// @Summary [FTName]PutHandler
// @Schemes
// @Description  [FTName]PutHandler
// @Tags [FTName]
// @Accept json
// @Produce json
// @Param Request-Id header string true "شناسه درخواست"
// @Param ورودی body [FTName|pascalcase]Request true "اطلاعات ورودی"
// @Param id path string true "شناسه"
// @Router /[FTName|kebabcase]/:id [put]
// @Security  OAuth2Password
// @Success 200 {object} [FTName|pascalcase]Response
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
func (env [FTName|camelcase]Env) [FTName|pascalcase]DeleteHandler(simulation bool) any {
	return handlers.BaseHandler[[FTName|pascalcase]Request, *[FTName|pascalcase]Response, [FTName|pascalcase]Handler](
		env.Interface, 
		[FTName|pascalcase]Handler{Name:"[FTName|kebabcase]-delete"}, 
		simulation,
	)
}


// [FTName]GetHandler godoc
// @Summary [FTName]GetHandler
// @Schemes
// @Description  [FTName]GetHandler
// @Tags [FTName]
// @Accept json
// @Produce json
// @Param Request-Id header string true "شناسه درخواست"
// @Param ورودی query [FTName|pascalcase]Row true "اطلاعات ورودی"
// @Router /[FTName|kebabcase] [get]
// @Security  OAuth2Password
// @Success 200 {object} [][FTName|pascalcase]Row
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
func (env [FTName|camelcase]Env) [FTName|pascalcase]GetHandler(simulation bool) any {
	return handlers.QueryHandler[[FTName|pascalcase]Row](
		"[FTName|kebabcase]-get", 
		QuerySingle, 
		"/[FTName|kebabcase]", 
		QueryMap, 
		env.Interface, 
		libRequest.QueryWithURI, 
		true, 
		simulation,
		nil,
	)
}


// [FTName]GetAllHandler godoc
// @Summary [FTName]GetAllHandler
// @Schemes
// @Description  [FTName]GetAllHandler
// @Tags [FTName]
// @Accept json
// @Produce json
// @Param Request-Id header string true "شناسه درخواست"
// @Router /[FTName|kebabcase]/all [get]
// @Security  OAuth2Password
// @Success 200 {object} [][FTName|pascalcase]Row
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
func (env [FTName|camelcase]Env) [FTName|pascalcase]GetAllHandler(simulation bool) any {
	return handlers.QueryHandler[[FTName|pascalcase]Row](
		"[FTName|kebabcase]-get-all", 
		QueryAll, 
		"/[FTName|kebabcase]/all", 
		QueryMap, 
		env.Interface, 
		libRequest.Query, 
		true, 
		simulation,
		nil,
	)
}
