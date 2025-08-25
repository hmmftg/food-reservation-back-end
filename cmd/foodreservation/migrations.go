package main

import (
	"log"

	"github.com/hmmftg/food-reservation-back-end/api/ums"
	"github.com/hmmftg/food-reservation-back-end/internal/params"
	"github.com/hmmftg/requestCore/libParams"
)

func (a Application) InitParams(wsParams *libParams.ApplicationParams[params.FoodReservationParams]) {
	dbParams := wsParams.GetDB(a.GetDbList()[0])

	err := ums.CreateTables(dbParams.Orm)

	if err != nil {
		log.Fatal(err)
	}
	ums.InsertDefaultRecords(dbParams.Orm)
}
