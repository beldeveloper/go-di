package main

import "fmt"

func main() {
	ctrl := InitializeController()
	// here we go
	report := ctrl.VotingReport()
	fmt.Println(report)
}
