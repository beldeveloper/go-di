package main

import (
	"github.com/beldeveloper/go-di/controller"
	"github.com/beldeveloper/go-di/repository"
	"github.com/beldeveloper/go-di/service"
	"go.uber.org/dig"
)

func PrepareContainer() *dig.Container {
	c := dig.New()
	providers := []interface{}{
		repository.NewUserMock,
		repository.NewTopicMock,
		service.NewVoteMock,
		controller.NewMock,
	}
	for _, p := range providers {
		err := c.Provide(p)
		if err != nil {
			panic(err)
		}
	}
	return c
}
