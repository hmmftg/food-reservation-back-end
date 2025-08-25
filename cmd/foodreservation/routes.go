package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hmmftg/food-reservation-back-end/api/ums"
	"github.com/hmmftg/food-reservation-back-end/internal/params"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libParams"
)

func (a Application) AddRoutes(
	model *requestCore.RequestCoreModel,
	wsParams *libParams.ApplicationParams[params.FoodReservationParams],
	roleMap map[string]string,
	rg *gin.RouterGroup,
) {
	api := rg.Group("api")
	ums.AddumsRoutes(model, wsParams, rg, api, false)

	rg.Static("/"+wsParams.Specific.StaticBaseUrl, wsParams.Specific.StaticPath)

	// Redirect root requests to '/ui'
	rg.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/"+wsParams.Specific.StaticBaseUrl)
	})

	a.engine.NoRoute(func(c *gin.Context) {
		// Catch-all route for React app
		if strings.HasPrefix(c.Request.RequestURI, "/"+wsParams.Specific.StaticBaseUrl) {
			c.File(wsParams.Specific.StaticPath + "/index.html")
		} else {
			c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": fmt.Sprintf("404 page %s not found", c.Request.RequestURI)})
		}
	})

}
