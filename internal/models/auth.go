package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	FullName string
	Email string
}

type Group struct {
	gorm.Model
	Name string
}

type Role struct {
	gorm.Model
	Name string
}
