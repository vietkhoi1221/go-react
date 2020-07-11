package actions

import (
	"fmt"
	"strconv"
	"model"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	db := Connect()
	users, _ := db.Select()
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	db := Connect()
	i, _ := strconv.Atoi(c.Param("id"))
	users, _ := db.SelectUser(i)
	c.JSON(200, users)
}

func AddUser(c *gin.Context) {
	db := Connect()
	var json model.User
	err := c.BindJSON(&json)
	if err == nil {
		err = db.Insert(json)
		if err == nil {
			c.JSON(200, "Successfully added user")
		} else {
			c.JSON(500, err)
		}
	}
	fmt.Println(err)
}

func DeleteUser(c *gin.Context) {
	db := Connect()
	i, _ := strconv.Atoi(c.Param("id"))
	err := db.DeleteUser(i)
	if err == nil {
		c.JSON(200, "Successfully deleted user")
	} else {
		c.JSON(500, err)
	}
}
