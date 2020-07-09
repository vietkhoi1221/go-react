package repository

import (
	models "model"
)

type UserRepository interface {
	Create() error
	Select() ([]models.User, error)
	Insert(u models.User) error
}
