package event

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

type SNSAlarm struct {
	*events.SNSEvent
}

func NewSNSAlarm(snsEvent *events.SNSEvent) (*SNSAlarm, error) {
	if snsEvent == nil {
		return nil, fmt.Errorf("snsEvent is nil")
	}
	if snsEvent.Records == nil {
		return nil, fmt.Errorf("snsEvent.Records is nil")
	}
	if len(snsEvent.Records) == 0 {
		return nil, fmt.Errorf("snsEvent.Records is empty")
	}

	return &SNSAlarm{
		SNSEvent: snsEvent,
	}, nil
}

func (s *SNSAlarm) GetAccountID() string {
	msg, err := s.getMessage()
	if err != nil {
		return "No Alarm Description"
	}

	return msg.AWSAccountID
}

func (s *SNSAlarm) GetSource() string {
	return s.Records[0].EventSource
}

func (s *SNSAlarm) GetAlarmName() string {
	msg, err := s.getMessage()
	if err != nil {
		return "No Alarm Name"
	}

	return msg.AlarmName
}

func (s *SNSAlarm) GetAlarmDescription() string {
	msg, err := s.getMessage()
	if err != nil {
		return "No Alarm Description"
	}

	return msg.AlarmDescription
}

func (s *SNSAlarm) getMessage() (*events.CloudWatchAlarmSNSPayload, error) {
	quotedMessage := s.SNSEvent.Records[0].SNS.Message

	var payload *events.CloudWatchAlarmSNSPayload
	if err := json.Unmarshal([]byte(quotedMessage), &payload); err != nil {
		log.Errorf("cannot unmarshal message: %s", err)
		return nil, err
	}
	return payload, nil
}
