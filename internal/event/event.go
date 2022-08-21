package event

type Eventer interface {
	GetSource() string
	GetAlarmName() string
}
