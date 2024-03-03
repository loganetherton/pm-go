package db

import (
	"fmt"
	"github.com/loganetherton/pm-go/config"
	"github.com/loganetherton/pm-go/db/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testProject models.TestProject
var testRelation models.TestRelation
var db *gorm.DB
var migrator gorm.Migrator
var err error

func Connect() {
	params := "charset=utf8mb4&parseTime=True&loc=Local"
	connect := fmt.Sprintf("%v:%v@tcp(%v)/%v", config.DbUser, config.DbPass, config.DbHost, config.DbName)
	dsn := connect + "?" + params
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate any table changes
	Migrate()

	//testProject := models.TestProject{
	//	Name:        "Test",
	//	Url:         "www.test.com",
	//	Description: "Test project",
	//	Another:     5,
	//}
	//res := db.Create(&testProject)
	//fmt.Println(res)

	//testProject = models.TestProject{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//}
	//queryRes := db.Find(&testProject)
	//if errors.Is(queryRes.Error, gorm.ErrRecordNotFound) {
	//	panic(gorm.ErrRecordNotFound)
	//}
	//fmt.Println(queryRes)
}

func Migrate() {
	if err = db.AutoMigrate(&testProject, &testRelation, &models.User{}); err != nil {
		panic(err)
	}
}

func DropTables() {
	migrator = db.Migrator()
	err = migrator.DropTable(&testProject, &testRelation)
	if err != nil {
		panic(err)
	}
}
