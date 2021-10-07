package repository

import "github.com/beldeveloper/go-di/model"

// NewTopicMock creates a new instance of the topic repository.
func NewTopicMock() Topic {
	return TopicMock{}
}

// TopicMock is a simple implementation of the topic repository.
type TopicMock struct {
}

// GetAll returns the hard-coded set of topics.
func (t TopicMock) GetAll() []model.Topic {
	return []model.Topic{
		model.Go,
		model.JavaScript,
		model.PHP,
	}
}
