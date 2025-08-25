package ums

import (
	"github.com/hmmftg/food-reservation-back-end/internal/params"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Role{},
		&User{},
	)
	if err != nil {
		return err
	}
	return nil
}

func InsertDefaultRecords(db *gorm.DB) {
	roles := []Role{
		{Model: params.Model{ID: "hmmftg", Name: "حمید"}},
		{Model: params.Model{ID: "majlesifar", Name: "پوریا"}},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles)
	users := []User{
		{Model: params.Model{ID: "hmmftg", Name: "حمید"}, Data: "", PersonID: "000014", Password: "059d0c8755f86b3b1ec4ab5790ec6516fca6b8182b73b845ee409d3ec8e155f9"},
		{Model: params.Model{ID: "majlesifar", Name: "پوریا"}, Data: "", PersonID: "000014", Password: "2ce09a479e179d41abf0bda9eaac25be5509c735944b74b22b082ba1b3ce8b6b"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)
}
