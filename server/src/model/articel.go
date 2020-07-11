package model

type Articel struct {
	ID     int    `form:"id" json:"id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}
