package model

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
