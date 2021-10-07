package repository

import "github.com/beldeveloper/go-di/model"

// User defines the interface of the user repository.
type User interface {
	GetAll() []model.User
}
