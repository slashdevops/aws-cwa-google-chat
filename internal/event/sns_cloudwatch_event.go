package event

import (
	"github.com/aws/aws-lambda-go/events"
)

type SNSCloudWatchEvent struct {
	*events.CloudWatchEvent
}

func NewSNSCloudWatchEvent(snsEvent *events.CloudWatchEvent) *SNSCloudWatchEvent {
	return &SNSCloudWatchEvent{
		CloudWatchEvent: snsEvent,
	}
}

func (s *SNSCloudWatchEvent) GetMessage() (*events.CloudWatchEvent, error) {
	return s.CloudWatchEvent, nil
}
