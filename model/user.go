package model

// Gender defines the type for storing user's gender.
type Gender string

const (
	// Male defines the male gender.
	Male Gender = "male"
	// Female defines the female gender.
	Female Gender = "female"
)

// User is a model of the user.
type User struct {
	Name   string
	Gender Gender
}
