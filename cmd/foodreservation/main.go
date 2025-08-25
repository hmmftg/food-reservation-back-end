package main

import (
	"fmt"
	"time"

	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hmmftg/food-reservation-back-end/api/ums"
	"github.com/hmmftg/food-reservation-back-end/internal/params"
	"github.com/hmmftg/requestCore"
	initiator "github.com/hmmftg/requestCore/libApplication"
	"github.com/hmmftg/requestCore/libParams"
	"github.com/swaggo/swag/v2"
)

type Application struct {
	engine *gin.Engine
}

var Version = "0.0.1"

func (a Application) Title() string {
	return "food-reservation API"
}
func (a Application) Name() string {
	return "food-reservation"
}
func (a Application) Version() string {
	return Version
}
func (a Application) BasePath() string {
	return ""
}
func (a Application) HasSwagger() bool {
	return false
}
func (a Application) SwaggerSpec() *swag.Spec {
	return nil
}
func (a Application) RequestFields() string {
	return ""
}
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

func (a *Application) InitGinApp(engine *gin.Engine) {
	a.engine = engine
}

func (a Application) GetCorsConfig() *cors.Config {
	return &cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Request-Id", "Branch-Id", "Person-Id", "User-Id", "X-Total-Count"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

func (a Application) GetDbList() []string {
	return []string{"food-reservation"}
}

func (a Application) GetKeys() [][]byte {
	keyByte := []byte{0xeb, 0xb2, 0x25, 0xcc, 0xe7, 0xfb, 0xa1, 0x5e, 0x32, 0xc6, 0xbb, 0xd0, 0xfd, 0x92, 0x05, 0x21}
	ivByte := []byte{0x4b, 0xdb, 0x3f, 0x59, 0xe9, 0x53, 0xb1, 0x16, 0xf2, 0x4d, 0xb0, 0xbe, 0xed, 0xcc, 0x12, 0x1d}
	return [][]byte{keyByte, ivByte}
}

func (a Application) InitParams(wsParams *libParams.ApplicationParams[params.FoodReservationParams]) {

}

// @title           سامانه سرویس های رزرو غذا
// @version         V0.0.1

// @description     مجموعه سرویس های مربوط به سامانه رزرو غذا

// @contact.name   پشتیبانی
// @contact.url    http://www.pooya.ir/support
// @contact.email  support@pooya.ir

// @servers.url https://reserve-food.pooya.ir/api
// @servers.description سرور رزرو غذا

// @securityDefinitions.bearerauth  احرازهویت-مشتری
// @bearerFormat: JWT
// @type: http
func main() {
	initiator.InitializeApp(&Application{})
}
