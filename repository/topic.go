package repository

import "github.com/beldeveloper/go-di/model"

// Topic defines the interface of the topic repository.
type Topic interface {
	GetAll() []model.Topic
}
