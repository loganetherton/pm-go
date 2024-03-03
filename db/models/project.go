package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Url         string
	Description string
}
