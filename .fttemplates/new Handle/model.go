package [FTName|camelcase?lowercase]

import (
	"main/lib/params"

	"github.com/hmmftg/requestCore"
	"github.com/hmmftg/requestCore/libQuery"
)

type [FTName|camelcase]Env struct {
	Params    params.ParamInterface
	Interface requestCore.RequestCoreInterface
}

type [FTName|pascalcase]Request struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type [FTName|pascalcase]Response struct {
	Result libQuery.DmlResult `json:"result"`
}

type [FTName|pascalcase]Handler struct {
	Name string
}

type [FTName|pascalcase]Row struct {
	ID string `form:"id" uri:"id" json:"id" db:"ID"`
}
const (
	QuerySingle = "single"
	QueryAll = "all"
)
var (
	QueryMap = map[string]libQuery.QueryCommand{
		QuerySingle: {
			Name:    QuerySingle,
			Command: "select * from SIMULATOR.[FTName|constantcase] where id = :1",
			Type:    libQuery.QuerySingle,
			Args:    []string{"id"},
		},
		QueryAll: {
			Name:    QueryAll,
			Command: "select * from SIMULATOR.[FTName|constantcase] order by id",
			Type:    libQuery.QueryAll,
		},
	}
)

func (env *[FTName|camelcase]Env) GetInterface() requestCore.RequestCoreInterface {
	return env.Interface
}
func (env *[FTName|camelcase]Env) GetParams() params.ParamInterface {
	return env.Params
}
func (env *[FTName|camelcase]Env) SetInterface(core requestCore.RequestCoreInterface) {
	env.Interface = core
}
func (env *[FTName|camelcase]Env) SetParams(parameters params.ParamInterface) {
	env.Params = parameters
}