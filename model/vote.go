package model

// Vote is a model that represents a vote of the particular user for the particular topic.
type Vote struct {
	UserName  string
	TopicName Topic
}
