package entities

import "gorm.io/gorm"

type User struct {
	gorm.DB
	name     string
	email    string
	password string
}
