package event

import (
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestSNSAlarm_NewSNSAlarm(t *testing.T) {
	t.Run("NewSNSAlarm", func(t *testing.T) {
		snsEvent := &events.SNSEvent{
			Records: []events.SNSEventRecord{
				{
					EventSource:          "aws:sns",
					EventVersion:         "1.0",
					EventSubscriptionArn: "arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
					SNS: events.SNSEntity{
						Type:              "Notification",
						MessageID:         "95df01b4-ee98-50d0-98c5-b5cee08f5279",
						TopicArn:          "arn:aws:sns:us-east-1:123456789012:my-topic",
						Subject:           "ALARM: \"Test LocalTime\" in US East (N. Virginia)",
						Message:           "{\"AlarmName\":\"Test LocalTime\",\"AlarmDescription\":\"Alarm Notification in my local timezone\",\"AWSAccountId\":\"[Account ID]\",\"NewStateValue\":\"ALARM\",\"NewStateReason\":\"Threshold Crossed: 1 out of the last 1 datapoints [0.0 (04/12/20 03:56:00)] was greater than or equal to the threshold (0.0) (minimum 1 datapoint for OK -> ALARM transition).\",\"StateChangeTime\":\"2020-12-04T03:57:01.659+0000\",\"Region\":\"US East (N. Virginia)\",\"AlarmArn\":\"arn:aws:cloudwatch:[region Id]:[Account ID]:alarm:Test LocalTime\",\"OldStateValue\":\"OK\",\"Trigger\":{\"Period\":60,\"EvaluationPeriods\":1,\"ComparisonOperator\":\"GreaterThanOrEqualToThreshold\",\"Threshold\":0.0,\"TreatMissingData\":\"- TreatMissingData:                    missing\",\"EvaluateLowSampleCountPercentile\":\"\",\"Metrics\":[{\"Expression\":\"FILL(m1, 0)\",\"Id\":\"e1\",\"Label\":\"Expression1\",\"ReturnData\":true},{\"Id\":\"m1\",\"MetricStat\":{\"Metric\":{\"Dimensions\":[{\"value\":\"API\",\"name\":\"Type\"},{\"value\":\"DescribeAlarms\",\"name\":\"Resource\"},{\"value\":\"CloudWatch\",\"name\":\"Service\"},{\"value\":\"None\",\"name\":\"Class\"}],\"MetricName\":\"CallCount\",\"Namespace\":\"AWS/Usage\"},\"Period\":60,\"Stat\":\"Average\"},\"ReturnData\":false}]}}",
						Timestamp:         time.Date(2020, 12, 0o4, 0o3, 57, 0o1, 659000000, time.UTC), //"2020-12-04T03:57:01.659Z"
						SignatureVersion:  "1",
						Signature:         "WcgVMPrlQsJY3yqbds968tqKPC6KKDWHSjIwEmzKVHZYg6foN9F5sm2Tp5IWPgaM9wMmYg8dpQjkxSm4q9V9iP1PbLp81RgJS2NghdeHNVnyxyzywXFMDztYZpgB2pjzfT101RVGpUwVPntOpBeBq2KAs/NrFX1nS2aTK/OX+gyOxwYZxRftzd+ttHA+PCh0kKlym7nnxaWuO9hgSrnupH2YttuvsdTSAOZ4MGhBON/sMmmlcxzfiFD+jJaqlHFmQ0DncjSe1NNwceOpwNsue6//sMYU1QzV6bO34I343KmQdXYw/KISDz7qH70Odm7nRLN3ExSOhtC/FS0/dXGl4Q==",
						SigningCertURL:    "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-bb750dd426d95ee9390147a5624348ee.pem",
						UnsubscribeURL:    "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
						MessageAttributes: make(map[string]interface{}),
					},
				},
			},
		}

		got, err := NewSNSAlarm(snsEvent)
		assert.NoError(t, err)
		assert.NotNil(t, got)

		want := &SNSAlarm{
			SNSEvent: snsEvent,
		}

		assert.Equal(t, want, got)
	})
}

func TestSNSAlarm_getMessage(t *testing.T) {
	t.Run("GetMessage", func(t *testing.T) {
		snsEvent := &events.SNSEvent{
			Records: []events.SNSEventRecord{
				{
					EventSource:          "aws:sns",
					EventVersion:         "1.0",
					EventSubscriptionArn: "arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
					SNS: events.SNSEntity{
						Type:              "Notification",
						MessageID:         "95df01b4-ee98-50d0-98c5-b5cee08f5279",
						TopicArn:          "arn:aws:sns:us-east-1:123456789012:my-topic",
						Subject:           "ALARM: \"Test LocalTime\" in US East (N. Virginia)",
						Message:           "{\"AlarmName\":\"Test LocalTime\",\"AlarmDescription\":\"Alarm Notification in my local timezone\",\"AWSAccountId\":\"[Account ID]\",\"NewStateValue\":\"ALARM\",\"NewStateReason\":\"Threshold Crossed: 1 out of the last 1 datapoints [0.0 (04/12/20 03:56:00)] was greater than or equal to the threshold (0.0) (minimum 1 datapoint for OK -> ALARM transition).\",\"StateChangeTime\":\"2020-12-04T03:57:01.659+0000\",\"Region\":\"US East (N. Virginia)\",\"AlarmArn\":\"arn:aws:cloudwatch:[region Id]:[Account ID]:alarm:Test LocalTime\",\"OldStateValue\":\"OK\",\"Trigger\":{\"Period\":60,\"EvaluationPeriods\":1,\"ComparisonOperator\":\"GreaterThanOrEqualToThreshold\",\"Threshold\":0.0,\"TreatMissingData\":\"- TreatMissingData:                    missing\",\"EvaluateLowSampleCountPercentile\":\"\",\"Metrics\":[{\"Expression\":\"FILL(m1, 0)\",\"Id\":\"e1\",\"Label\":\"Expression1\",\"ReturnData\":true},{\"Id\":\"m1\",\"MetricStat\":{\"Metric\":{\"Dimensions\":[{\"value\":\"API\",\"name\":\"Type\"},{\"value\":\"DescribeAlarms\",\"name\":\"Resource\"},{\"value\":\"CloudWatch\",\"name\":\"Service\"},{\"value\":\"None\",\"name\":\"Class\"}],\"MetricName\":\"CallCount\",\"Namespace\":\"AWS/Usage\"},\"Period\":60,\"Stat\":\"Average\"},\"ReturnData\":false}]}}",
						Timestamp:         time.Date(2020, 12, 0o4, 0o3, 57, 0o1, 659000000, time.UTC), //"2020-12-04T03:57:01.659Z"
						SignatureVersion:  "1",
						Signature:         "WcgVMPrlQsJY3yqbds968tqKPC6KKDWHSjIwEmzKVHZYg6foN9F5sm2Tp5IWPgaM9wMmYg8dpQjkxSm4q9V9iP1PbLp81RgJS2NghdeHNVnyxyzywXFMDztYZpgB2pjzfT101RVGpUwVPntOpBeBq2KAs/NrFX1nS2aTK/OX+gyOxwYZxRftzd+ttHA+PCh0kKlym7nnxaWuO9hgSrnupH2YttuvsdTSAOZ4MGhBON/sMmmlcxzfiFD+jJaqlHFmQ0DncjSe1NNwceOpwNsue6//sMYU1QzV6bO34I343KmQdXYw/KISDz7qH70Odm7nRLN3ExSOhtC/FS0/dXGl4Q==",
						SigningCertURL:    "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-bb750dd426d95ee9390147a5624348ee.pem",
						UnsubscribeURL:    "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
						MessageAttributes: make(map[string]interface{}),
					},
				},
			},
		}

		want := &events.CloudWatchAlarmSNSPayload{
			AlarmName:        "Test LocalTime",
			AlarmDescription: "Alarm Notification in my local timezone",
			AWSAccountID:     "[Account ID]",
			NewStateValue:    "ALARM",
			NewStateReason:   "Threshold Crossed: 1 out of the last 1 datapoints [0.0 (04/12/20 03:56:00)] was greater than or equal to the threshold (0.0) (minimum 1 datapoint for OK -> ALARM transition).",
			StateChangeTime:  "2020-12-04T03:57:01.659+0000",
			Region:           "US East (N. Virginia)",
			AlarmARN:         "arn:aws:cloudwatch:[region Id]:[Account ID]:alarm:Test LocalTime",
			OldStateValue:    "OK",
			Trigger: events.CloudWatchAlarmTrigger{
				Period:                           60,
				EvaluationPeriods:                1,
				ComparisonOperator:               "GreaterThanOrEqualToThreshold",
				Threshold:                        0,
				TreatMissingData:                 "- TreatMissingData:                    missing",
				EvaluateLowSampleCountPercentile: "",
				Metrics: []events.CloudWatchMetricDataQuery{
					{
						Expression: "FILL(m1, 0)",
						ID:         "e1",
						Label:      "Expression1",
						ReturnData: true,
					},
					{
						ID:         "m1",
						ReturnData: false,
						MetricStat: events.CloudWatchMetricStat{
							Metric: events.CloudWatchMetric{
								Dimensions: []events.CloudWatchDimension{
									{
										Value: "API",
										Name:  "Type",
									},
									{
										Value: "DescribeAlarms",
										Name:  "Resource",
									},
									{
										Value: "CloudWatch",
										Name:  "Service",
									},
									{
										Value: "None",
										Name:  "Class",
									},
								},
								MetricName: "CallCount",
								Namespace:  "AWS/Usage",
							},
							Period: 60,
							Stat:   "Average",
						},
					},
				},
			},
		}

		s, err := NewSNSAlarm(snsEvent)
		assert.NoError(t, err)
		assert.NotNil(t, s)

		got, err := s.getMessage()
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}

func TestSNSAlarm_GetSource(t *testing.T) {
	t.Run("Full SNSEvent", func(t *testing.T) {
		snsEvent := &events.SNSEvent{
			Records: []events.SNSEventRecord{
				{
					EventSource:          "aws:sns",
					EventVersion:         "1.0",
					EventSubscriptionArn: "arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
					SNS: events.SNSEntity{
						Type:              "Notification",
						MessageID:         "95df01b4-ee98-50d0-98c5-b5cee08f5279",
						TopicArn:          "arn:aws:sns:us-east-1:123456789012:my-topic",
						Subject:           "ALARM: \"Test LocalTime\" in US East (N. Virginia)",
						Message:           "{\"AlarmName\":\"Test LocalTime\",\"AlarmDescription\":\"Alarm Notification in my local timezone\",\"AWSAccountId\":\"[Account ID]\",\"NewStateValue\":\"ALARM\",\"NewStateReason\":\"Threshold Crossed: 1 out of the last 1 datapoints [0.0 (04/12/20 03:56:00)] was greater than or equal to the threshold (0.0) (minimum 1 datapoint for OK -> ALARM transition).\",\"StateChangeTime\":\"2020-12-04T03:57:01.659+0000\",\"Region\":\"US East (N. Virginia)\",\"AlarmArn\":\"arn:aws:cloudwatch:[region Id]:[Account ID]:alarm:Test LocalTime\",\"OldStateValue\":\"OK\",\"Trigger\":{\"Period\":60,\"EvaluationPeriods\":1,\"ComparisonOperator\":\"GreaterThanOrEqualToThreshold\",\"Threshold\":0.0,\"TreatMissingData\":\"- TreatMissingData:                    missing\",\"EvaluateLowSampleCountPercentile\":\"\",\"Metrics\":[{\"Expression\":\"FILL(m1, 0)\",\"Id\":\"e1\",\"Label\":\"Expression1\",\"ReturnData\":true},{\"Id\":\"m1\",\"MetricStat\":{\"Metric\":{\"Dimensions\":[{\"value\":\"API\",\"name\":\"Type\"},{\"value\":\"DescribeAlarms\",\"name\":\"Resource\"},{\"value\":\"CloudWatch\",\"name\":\"Service\"},{\"value\":\"None\",\"name\":\"Class\"}],\"MetricName\":\"CallCount\",\"Namespace\":\"AWS/Usage\"},\"Period\":60,\"Stat\":\"Average\"},\"ReturnData\":false}]}}",
						Timestamp:         time.Date(2020, 12, 0o4, 0o3, 57, 0o1, 659000000, time.UTC), //"2020-12-04T03:57:01.659Z"
						SignatureVersion:  "1",
						Signature:         "WcgVMPrlQsJY3yqbds968tqKPC6KKDWHSjIwEmzKVHZYg6foN9F5sm2Tp5IWPgaM9wMmYg8dpQjkxSm4q9V9iP1PbLp81RgJS2NghdeHNVnyxyzywXFMDztYZpgB2pjzfT101RVGpUwVPntOpBeBq2KAs/NrFX1nS2aTK/OX+gyOxwYZxRftzd+ttHA+PCh0kKlym7nnxaWuO9hgSrnupH2YttuvsdTSAOZ4MGhBON/sMmmlcxzfiFD+jJaqlHFmQ0DncjSe1NNwceOpwNsue6//sMYU1QzV6bO34I343KmQdXYw/KISDz7qH70Odm7nRLN3ExSOhtC/FS0/dXGl4Q==",
						SigningCertURL:    "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-bb750dd426d95ee9390147a5624348ee.pem",
						UnsubscribeURL:    "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
						MessageAttributes: make(map[string]interface{}),
					},
				},
			},
		}
		s, err := NewSNSAlarm(snsEvent)
		assert.NoError(t, err)
		assert.NotNil(t, s)

		got := s.GetSource()
		assert.NotNil(t, got)
		assert.Equal(t, "aws:sns", got)
	})
}

func TestSNSAlarm_GetAlarmName(t *testing.T) {
	t.Run("Full SNSEvent", func(t *testing.T) {
		snsEvent := &events.SNSEvent{
			Records: []events.SNSEventRecord{
				{
					EventSource:          "aws:sns",
					EventVersion:         "1.0",
					EventSubscriptionArn: "arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
					SNS: events.SNSEntity{
						Type:              "Notification",
						MessageID:         "95df01b4-ee98-50d0-98c5-b5cee08f5279",
						TopicArn:          "arn:aws:sns:us-east-1:123456789012:my-topic",
						Subject:           "ALARM: \"Test LocalTime\" in US East (N. Virginia)",
						Message:           "{\"AlarmName\":\"Test LocalTime\",\"AlarmDescription\":\"Alarm Notification in my local timezone\",\"AWSAccountId\":\"[Account ID]\",\"NewStateValue\":\"ALARM\",\"NewStateReason\":\"Threshold Crossed: 1 out of the last 1 datapoints [0.0 (04/12/20 03:56:00)] was greater than or equal to the threshold (0.0) (minimum 1 datapoint for OK -> ALARM transition).\",\"StateChangeTime\":\"2020-12-04T03:57:01.659+0000\",\"Region\":\"US East (N. Virginia)\",\"AlarmArn\":\"arn:aws:cloudwatch:[region Id]:[Account ID]:alarm:Test LocalTime\",\"OldStateValue\":\"OK\",\"Trigger\":{\"Period\":60,\"EvaluationPeriods\":1,\"ComparisonOperator\":\"GreaterThanOrEqualToThreshold\",\"Threshold\":0.0,\"TreatMissingData\":\"- TreatMissingData:                    missing\",\"EvaluateLowSampleCountPercentile\":\"\",\"Metrics\":[{\"Expression\":\"FILL(m1, 0)\",\"Id\":\"e1\",\"Label\":\"Expression1\",\"ReturnData\":true},{\"Id\":\"m1\",\"MetricStat\":{\"Metric\":{\"Dimensions\":[{\"value\":\"API\",\"name\":\"Type\"},{\"value\":\"DescribeAlarms\",\"name\":\"Resource\"},{\"value\":\"CloudWatch\",\"name\":\"Service\"},{\"value\":\"None\",\"name\":\"Class\"}],\"MetricName\":\"CallCount\",\"Namespace\":\"AWS/Usage\"},\"Period\":60,\"Stat\":\"Average\"},\"ReturnData\":false}]}}",
						Timestamp:         time.Date(2020, 12, 0o4, 0o3, 57, 0o1, 659000000, time.UTC), //"2020-12-04T03:57:01.659Z"
						SignatureVersion:  "1",
						Signature:         "WcgVMPrlQsJY3yqbds968tqKPC6KKDWHSjIwEmzKVHZYg6foN9F5sm2Tp5IWPgaM9wMmYg8dpQjkxSm4q9V9iP1PbLp81RgJS2NghdeHNVnyxyzywXFMDztYZpgB2pjzfT101RVGpUwVPntOpBeBq2KAs/NrFX1nS2aTK/OX+gyOxwYZxRftzd+ttHA+PCh0kKlym7nnxaWuO9hgSrnupH2YttuvsdTSAOZ4MGhBON/sMmmlcxzfiFD+jJaqlHFmQ0DncjSe1NNwceOpwNsue6//sMYU1QzV6bO34I343KmQdXYw/KISDz7qH70Odm7nRLN3ExSOhtC/FS0/dXGl4Q==",
						SigningCertURL:    "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-bb750dd426d95ee9390147a5624348ee.pem",
						UnsubscribeURL:    "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:123456789012:my-topic:80289ba6-0fd4-4079-afb4-ce8c8260f0ca",
						MessageAttributes: make(map[string]interface{}),
					},
				},
			},
		}
		s, err := NewSNSAlarm(snsEvent)
		assert.NoError(t, err)
		assert.NotNil(t, s)

		got := s.GetAlarmName()
		assert.NotNil(t, got)
		assert.Equal(t, "Test LocalTime", got)
	})
}
