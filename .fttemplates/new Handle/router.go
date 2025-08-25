package [FTName|camelcase?lowercase]

import (
	"main/lib/params"

	"github.com/gin-gonic/gin"
	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libGin"
)

func Add[FTName|pascalcase]Routes(
	model *requestCore.RequestCoreModel,
	wsParams *params.CardIssueWsParams,
	_ map[string]string,
	rg *gin.RouterGroup,
	simulation bool,
) {
	env := &[FTName|camelcase]Env{
		Interface: model,
		Params:    wsParams,
	}
	root := rg.Group("/[FTName|kebabcase]")
	root.GET("all", libGin.Gin(env.[FTName|pascalcase]GetAllHandler(simulation)))
	root.GET(":id", libGin.Gin(env.[FTName|pascalcase]GetHandler(simulation)))
	root.POST("", libGin.Gin(env.[FTName|pascalcase]PostHandler(simulation)))
	root.PUT(":id", libGin.Gin(env.[FTName|pascalcase]PutHandler(simulation)))
	root.DELETE(":id", libGin.Gin(env.[FTName|pascalcase]DeleteHandler(simulation)))
}
