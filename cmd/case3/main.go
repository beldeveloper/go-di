package main

import (
	"fmt"
	"github.com/beldeveloper/go-di/controller"
)

func main() {
	c := PrepareContainer()
	err := c.Invoke(func(ctrl controller.Controller) {
		// here we go
		report := ctrl.VotingReport()
		fmt.Println(report)
	})
	if err != nil {
		panic(err)
	}
}
