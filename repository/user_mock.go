package repository

import "github.com/beldeveloper/go-di/model"

// NewUserMock creates a new instance of the user repository.
func NewUserMock() User {
	return UserMock{}
}

// UserMock is a simple implementation of the user repository.
type UserMock struct {
}

// GetAll returns the hard-coded set of users.
func (u UserMock) GetAll() []model.User {
	return []model.User{
		{
			Name:   "John",
			Gender: model.Male,
		},
		{
			Name:   "Alice",
			Gender: model.Female,
		},
		{
			Name:   "Kate",
			Gender: model.Female,
		},
	}
}
