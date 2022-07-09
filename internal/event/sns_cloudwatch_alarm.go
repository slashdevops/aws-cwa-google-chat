package event

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
)

type SNSAlarm struct {
	*events.SNSEvent
}

func NewSNSAlarm(snsEvent *events.SNSEvent) *SNSAlarm {
	return &SNSAlarm{
		SNSEvent: snsEvent,
	}
}

func (s *SNSAlarm) GetMessage() (*events.CloudWatchAlarmSNSPayload, error) {
	quotedMessage := s.SNSEvent.Records[0].SNS.Message
	stringMessage, err := strconv.Unquote(quotedMessage)
	if err != nil {
		log.Errorf("cannot unquote message: %s", err)
		return nil, err
	}

	var payload *events.CloudWatchAlarmSNSPayload
	if err := json.Unmarshal([]byte(stringMessage), &payload); err != nil {
		log.Errorf("cannot unmarshal message: %s", err)
		return nil, err
	}
	return payload, nil
}
