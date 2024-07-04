package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"unique;not null"`
	Age  int64  `json:"age" binding:"required" gorm:"unique;not null"`
}
