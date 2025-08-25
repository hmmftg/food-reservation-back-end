package ums

import (
	"github.com/gin-gonic/gin"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libGin"
	"github.com/hmmftg/requestCore/libParams"
)

func AddumsRoutes(
	model *requestCore.RequestCoreModel,
	wsParams libParams.ParamInterface,
	rg *gin.RouterGroup,
	api *gin.RouterGroup,
	simulation bool,
) {
	env := &umsEnv{
		Interface: model,
		Params:    wsParams,
	}
	root := rg.Group("/ums")
	root.POST("/auth/login/", libGin.Gin(env.umsLogin(simulation)))
	root.POST("/register/", libGin.Gin(env.umsRegister(simulation)))
	root.PUT("/logout/", libGin.Gin(env.umsLogout(simulation)))
	api.Use(libGin.Gin(env.UmsIntrospect("service auth middleware", ServiceAuthHandler{})))
	rootApi := api.Group("/ums")
	rootApi.GET("/check/", libGin.Gin(env.umsCheck(simulation)))
	rootApi.GET("/permissions/", libGin.Gin(env.umsPermissions(simulation)))
	rootApi.GET("/user/", libGin.Gin(env.umsGetUser(simulation)))
}
