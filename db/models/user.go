package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Role      int
	Active    bool
	Password  string
	LastLogin time.Time
}
