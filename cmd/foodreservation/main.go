package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hmmftg/food-reservation-back-end/internal/params"
	initiator "github.com/hmmftg/requestCore/libApplication"
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
	app := initiator.InitializeApp[params.FoodReservationParams](&Application{})
	initiator.StartApp(*app)
}
