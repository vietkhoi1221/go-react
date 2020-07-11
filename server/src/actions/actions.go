package actions

import (
	"driver"
	repo "repository"
	implement "repository/implement"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "myuser"
	password = "123456"
	dbname   = "mydb"
)

func Connect() repo.UserRepository {
	db := driver.Connect(host, port, user, password, dbname)
	error := db.SQL.Ping()
	if error != nil {
		panic(error)
	}
	userRepo := implement.UserRepo(db.SQL)
	return userRepo
}
