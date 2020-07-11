package repository

import (
	models "model"
)

type UserRepository interface {
	CreateUser() error
	SelectUsers() ([]models.User, error)
	SelectUser(id int) ([]models.User, error)
	InsertUser(u models.User) error
	DeleteUser(id int) error
}

type ArticleRepository interface {
	CreateArticle() error
	SelectArticle(id int) ([]models.Article, error)
	SelectArticle(id int) ([]models.Article, error)
	InsertArticle(a models.Article) error
	DeleteArticle(id int) error
}
