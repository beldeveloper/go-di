package service

import (
	"github.com/beldeveloper/go-di/model"
	"github.com/beldeveloper/go-di/repository"
)

// NewVoteMock creates a new instance of the vote service.
func NewVoteMock(users repository.User, topics repository.Topic) Vote {
	return VoteMock{users: users, topics: topics}
}

// VoteMock is a simple implementation of the vote service.
type VoteMock struct {
	users  repository.User
	topics repository.Topic
}

// PerformVoting returns the voting result of all users.
func (s VoteMock) PerformVoting() []model.Vote {
	users := s.users.GetAll()
	topics := s.topics.GetAll()
	votes := make([]model.Vote, 0, len(users))
	for _, u := range users {
		for _, t := range topics {
			if !s.wantsVote(u, t) {
				continue
			}
			votes = append(votes, model.Vote{
				UserName:  u.Name,
				TopicName: t,
			})
			break
		}
	}
	return votes
}

// wantsVote returns true if the particular user wants to vote for the particular topic.
func (s VoteMock) wantsVote(u model.User, t model.Topic) bool {
	switch u.Gender {
	case model.Male:
		return t == model.Go
	case model.Female:
		return t == model.JavaScript
	default:
		return t == model.PHP
	}
}
