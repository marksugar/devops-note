package config

import "time"

type HuaweicloudSlowPostMessage struct {
	Message string `json:"message"`
}
type LogConfig struct {
	Filename     string
	MaxSize      int
	MaxBackups   int
	LocalTime    bool
	Compress     bool
	Level        int
	ReportCaller bool
}

type ExporterConfig struct {
	Log LogConfig
}

type HuaweicloudSlow struct {
	Signature        string      `json:"signature"`
	Subject          string      `json:"subject"`
	TopicUrn         string      `json:"topic_urn"`
	MessageID        string      `json:"message_id"`
	SignatureVersion string      `json:"signature_version"`
	Type             string      `json:"type"`
	Message          interface{} `json:"message"`
	UnsubscribeURL   string      `json:"unsubscribe_url"`
	SigningCertURL   string      `json:"signing_cert_url"`
	Timestamp        time.Time   `json:"timestamp"`
}

type HuaweicloudSlowMessage struct {
	MessageType        string `json:"message_type"`
	AlarmID            string `json:"alarm_id"`
	AlarmName          string `json:"alarm_name"`
	AlarmStatus        string `json:"alarm_status"`
	Time               int64  `json:"time"`
	Namespace          string `json:"namespace"`
	MetricName         string `json:"metric_name"`
	Dimension          string `json:"dimension"`
	Period             int    `json:"period"`
	Filter             string `json:"filter"`
	ComparisonOperator string `json:"comparison_operator"`
	Value              int    `json:"value"`
	Unit               string `json:"unit"`
	Count              int    `json:"count"`
	AlarmValue         []struct {
		Time  int `json:"time"`
		Value int `json:"value"`
	} `json:"alarmValue"`
	SmsContent       string `json:"sms_content"`
	DefaultContent   string `json:"default_content"`
	TemplateVariable struct {
		AccountName    string `json:"AccountName"`
		Namespace      string `json:"Namespace"`
		DimensionName  string `json:"DimensionName"`
		ResourceName   string `json:"ResourceName"`
		MetricName     string `json:"MetricName"`
		IsAlarm        bool   `json:"IsAlarm"`
		IsCycleTrigger bool   `json:"IsCycleTrigger"`
		AlarmLevel     string `json:"AlarmLevel"`
		Region         string `json:"Region"`
		ResourceID     string `json:"ResourceId"`
		PrivateIP      string `json:"PrivateIp"`
		AlarmRule      string `json:"AlarmRule"`
		CurrentData    string `json:"CurrentData"`
		AlarmTime      string `json:"AlarmTime"`
		DataPoint      struct {
			Two0230822094600GMT0800 string `json:"2023/08/22 09:46:00 GMT+08:00"`
		} `json:"DataPoint"`
		DataPointTime             []string `json:"DataPointTime"`
		AlarmRuleName             string   `json:"AlarmRuleName"`
		AlarmID                   string   `json:"AlarmId"`
		AlarmDesc                 string   `json:"AlarmDesc"`
		MonitoringRange           string   `json:"MonitoringRange"`
		IsOriginalValue           bool     `json:"IsOriginalValue"`
		Period                    string   `json:"Period"`
		Filter                    string   `json:"Filter"`
		ComparisonOperator        string   `json:"ComparisonOperator"`
		Value                     string   `json:"Value"`
		Unit                      string   `json:"Unit"`
		Count                     int      `json:"Count"`
		EventContent              string   `json:"EventContent"`
		Link                      string   `json:"Link"`
		EpName                    string   `json:"EpName"`
		IsIEC                     bool     `json:"IsIEC"`
		IsAgentEvent              bool     `json:"IsAgentEvent"`
		IngressMaxBandwidthPerSec string   `json:"IngressMaxBandwidthPerSec"`
		EgressMaxBandwidthPerSec  string   `json:"EgressMaxBandwidthPerSec"`
	} `json:"template_variable"`
}
