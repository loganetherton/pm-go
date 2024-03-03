package models

import "gorm.io/gorm"

type TestProject struct {
	gorm.Model
	Name        string
	Url         string
	Description string
}

type TestRelation struct {
	gorm.Model
	Project   TestProject
	ProjectId int
	Name      string
}
