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

func (s *SNSCloudWatchEvent) GetSource() string {
	return "CloudWatchEvent"
}

func (s *SNSCloudWatchEvent) GetAlarmName() string {
	msg, err := s.GetMessage()
	if err != nil {
		return ""
	}

	return msg.ID
}

func (s *SNSCloudWatchEvent) GetMessage() (*events.CloudWatchEvent, error) {
	return s.CloudWatchEvent, nil
}
