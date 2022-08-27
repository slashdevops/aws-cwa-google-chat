package event

type Eventer interface {
	GetAccountID() string
	GetSource() string
	GetAlarmName() string
	GetAlarmDescription() string
}
