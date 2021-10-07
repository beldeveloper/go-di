package service

import "github.com/beldeveloper/go-di/model"

// Vote defines the interface of the vote service.
type Vote interface {
	PerformVoting() []model.Vote
}
