package models

import "gorm.io/gorm"

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
}
