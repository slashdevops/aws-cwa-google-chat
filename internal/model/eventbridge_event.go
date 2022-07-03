package model

import "time"

// Reference: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch-and-eventbridge.html

type AWSCWAEvent struct {
	Version    string            `json:"version"`
	ID         string            `json:"id"`
	DetailType string            `json:"detail-type"`
	Source     string            `json:"source"`
	Account    string            `json:"account"`
	Time       time.Time         `json:"time"`
	Region     string            `json:"region"`
	Resources  []string          `json:"resources"`
	Detail     AWSCWAEventDetail `json:"detail"`
}

type AWSCWAEventDetail struct {
	AlarmName             string                         `json:"alarmName"`
	Operation             string                         `json:"operation"`
	Configuration         AWSCWAEventDetailConfiguration `json:"configuration"`
	PreviousConfiguration AWSCWAEventDetailConfiguration `json:"previousConfiguration"`
	PreviousState         AWSCWAEventDetailState         `json:"previousState"`
	State                 AWSCWAEventDetailState         `json:"state"`
}

type AWSCWAEventDetailConfiguration struct {
	AlarmName               string                                 `json:"alarmName"`
	AlarmRule               string                                 `json:"alarmRule"`
	AlarmActions            []string                               `json:"alarmActions"`
	ActionsEnabled          bool                                   `json:"actionsEnabled"`
	ComparisonOperator      string                                 `json:"comparisonOperator"`
	EvaluationPeriods       int                                    `json:"evaluationPeriods"`
	InsufficientDataActions []string                               `json:"insufficientDataActions"`
	Description             string                                 `json:"description"`
	OKActions               []string                               `json:"okActions"`
	Metrics                 []AWSCWAEventDetailConfigurationMetric `json:"metrics"`
	Timestamp               time.Time                              `json:"timestamp"`
	Threshold               float32                                `json:"threshold"`
	TreatMissingData        string                                 `json:"treatMissingData"`
}

type AWSCWAEventDetailState struct {
	Reason     string            `json:"reason"`
	ReasonData map[string]string `json:"reasonData"`
	Timestamp  time.Time         `json:"timestamp"`
	Value      string            `json:"value"`
}

type AWSCWAEventDetailConfigurationMetric struct {
	ID         string                                   `json:"id"`
	Expression string                                   `json:"expression"`
	MetricStat AWSCWAEventDetailConfigurationMetricStat `json:"metricStat"`
	Label      string                                   `json:"label"`
	ReturnData bool                                     `json:"returnData"`
}

type AWSCWAEventDetailConfigurationMetricStat struct {
	Period int                                            `json:"period"`
	Stat   string                                         `json:"stat"`
	Metric AWSCWAEventDetailConfigurationMetricStatMetric `json:"metric"`
}

type AWSCWAEventDetailConfigurationMetricStatMetric struct {
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	Dimensions map[string]string `json:"dimensions"`
}
