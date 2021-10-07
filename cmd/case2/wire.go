//+build wireinject

package main

import (
	"github.com/beldeveloper/go-di/controller"
	"github.com/beldeveloper/go-di/repository"
	"github.com/beldeveloper/go-di/service"
	"github.com/google/wire"
)

func InitializeController() controller.Controller {
	wire.Build(
		controller.NewMock,
		service.NewVoteMock,
		repository.NewUserMock,
		repository.NewTopicMock,
	)
	return nil
}
