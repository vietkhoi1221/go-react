package model

type User struct {
	ID     int    `form:"id" json:"id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
	Gender string `form:"gender" json:"gender" binding:"required"`
	Email  string `form:"email" json:"email" binding:"required"`
}
