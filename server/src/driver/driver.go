package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	SQL *sql.DB
}

var Postgres = &PostgreSQL{}

func Connect(host, port, user, password, dbname string) *PostgreSQL {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	Postgres.SQL = db
	return Postgres
}
