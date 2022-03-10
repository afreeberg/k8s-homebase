package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	FullName string
	Email string
}
