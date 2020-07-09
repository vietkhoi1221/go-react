package implement

import (
	"database/sql"
	"fmt"
	models "model"
	repo "repository"
)

type UserRepositoryImplement struct {
	Db *sql.DB
}

func UserRepo(db *sql.DB) repo.UserRepository {
	return &UserRepositoryImplement{Db: db}
}

func (u *UserRepositoryImplement) Create() error {
	sqlQuery := `CREATE TABLE users (ID integer CONSTRAINT firstkey PRIMARY KEY, Name varchar(40) NOT NULL, Gender varchar(40) NOT NULL, Email varchar(40) NOT NULL)`
	_, error := u.Db.Exec(sqlQuery)
	fmt.Println(error)
	return error
}

func (u *UserRepositoryImplement) Select() ([]models.User, error) {
	users := make([]models.User, 0)
	rows, error := u.Db.Query("SELECT ID, Name, Gender, Email FROM users")
	fmt.Println(error)
	if error != nil {
		return users, error
	}
	for rows.Next() {
		// var id int
		// var name, gender, email string
		user := models.User{}
		// error := rows.Scan(&id, &name, &gender, &email)
		error := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if error != nil {
			break
		}
		// user.ID = id
		// user.Name = name
		// user.Gender = gender
		// user.Email = email
		users = append(users, user)
	}
	error = rows.Err()
	if error != nil {
		return users, error
	}
	return users, nil
}

func (u *UserRepositoryImplement) Insert(user models.User) error {
	insertSQL := `INSERT INTO users (id, name, gender, email)
	VALUES ($1, $2, $3, $4)`
	_, error := u.Db.Exec(insertSQL, user.ID, user.Name, user.Gender, user.Email)
	return error
}
