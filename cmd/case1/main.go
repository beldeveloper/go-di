package main

import (
	"fmt"
	"github.com/beldeveloper/go-di/controller"
	"github.com/beldeveloper/go-di/repository"
	"github.com/beldeveloper/go-di/service"
)

func main() {
	users := repository.NewUserMock()
	topic := repository.NewTopicMock()
	votes := service.NewVoteMock(users, topic)
	ctrl := controller.NewMock(votes)
	// here we go
	report := ctrl.VotingReport()
	fmt.Println(report)
}
