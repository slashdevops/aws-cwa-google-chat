package event

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type SNSCloudWatchEvent struct {
	*events.CloudWatchEvent
}

func NewSNSCloudWatchEvent(snsEvent *events.CloudWatchEvent) (*SNSCloudWatchEvent, error) {
	if snsEvent == nil {
		return nil, fmt.Errorf("snsEvent is nil")
	}
	if snsEvent.Detail == nil {
		return nil, fmt.Errorf("snsEvent.Detail is nil")
	}

	return &SNSCloudWatchEvent{
		CloudWatchEvent: snsEvent,
	}, nil
}

func (s *SNSCloudWatchEvent) GetAccountID() string {
	return s.AccountID
}

func (s *SNSCloudWatchEvent) GetSource() string {
	return s.Source
}

func (s *SNSCloudWatchEvent) GetAlarmName() string {
	// msg, err := s.getMessage()
	// if err != nil {
	// 	return ""
	// }

	// return msg.Detail
	return ""
}

func (s *SNSCloudWatchEvent) GetAlarmDescription() string {
	return s.DetailType
}

func (s *SNSCloudWatchEvent) getMessage() (*events.CloudWatchEvent, error) {
	return s.CloudWatchEvent, nil
}

func (s *SNSCloudWatchEvent) getMessageDetail() string {
	// var payload *events.CloudWatchAlarmSNSPayload
	// if err := json.Unmarshal([]byte(quotedMessage), &payload); err != nil {
	// 	log.Errorf("cannot unmarshal message: %s", err)
	// 	return nil, err
	// }
	return ""
}
