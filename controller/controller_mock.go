package controller

import (
	"fmt"
	"github.com/beldeveloper/go-di/service"
)

// NewMock expects a votes service and creates a new instance of the controller.
func NewMock(votes service.Vote) Controller {
	return Mock{votes: votes}
}

// Mock is a simple implementation of the controller.
type Mock struct {
	votes service.Vote
}

// VotingReport performs the voting and collects the result in one string.
func (m Mock) VotingReport() string {
	votes := m.votes.PerformVoting()
	var report string
	for _, v := range votes {
		report += fmt.Sprintf("%s votes for %s\n", v.UserName, v.TopicName)
	}
	return report
}
