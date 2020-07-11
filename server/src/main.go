package main

import (
	"github.com/gin-gonic/gin"
	action "actions"
)

func main() {
	r := gin.Default()
	r.GET("/api/users", action.GetAllUser)
	r.GET("/api/user/:id", action.GetUser)
	r.POST("/api/user/create", action.AddUser)
	r.DELETE("/api/user/:id", action.DeleteUser)
	r.Run(":6005")
}
