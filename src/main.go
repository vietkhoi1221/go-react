package main

import (
	"driver"
	repo "repository"
	implement "repository/implement"

	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "myuser"
	password = "123456"
	dbname   = "mydb"
)

func connect() repo.UserRepository {
	db := driver.Connect(host, port, user, password, dbname)
	error := db.SQL.Ping()
	if error != nil {
		panic(error)
	}
	userRepo := implement.UserRepo(db.SQL)
	return userRepo
}

func getAllUser(c *gin.Context) {
	db := connect()
	users, _ := db.Select()
	c.JSON(200, users)
}

func getUser(c *gin.Context) {
	db := connect()
	users, _ := db.Select()
	c.JSON(200, users)
}

func main() {
	r := gin.Default()
	// r.GET("", Welcome)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	r.GET("/api/users", getAllUser)
	r.GET("/api/users/:id", getUser)
	r.Run(":8888")
}
