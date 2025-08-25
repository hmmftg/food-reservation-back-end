package ums

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID       string `gorm:"primarykey"`
	UserName string
	UserData string
	PersonID string
	Pass     string
}

func CreateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
	)
	if err != nil {
		return err
	}
	return nil
}

func InsertDefaultRecords(db *gorm.DB) {
	users := []User{
		{ID: "hmmftg", UserName: "حمید", UserData: "", PersonID: "000014", Pass: "059d0c8755f86b3b1ec4ab5790ec6516fca6b8182b73b845ee409d3ec8e155f9"},
		{ID: "majlesifar", UserName: "پوریا", UserData: "", PersonID: "000014", Pass: "2ce09a479e179d41abf0bda9eaac25be5509c735944b74b22b082ba1b3ce8b6b"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)
}
