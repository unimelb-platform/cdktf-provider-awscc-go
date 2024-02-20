package cloudwatchalarm


type CloudwatchAlarmMetricsMetricStat struct {
	// The metric to return, including the metric name, namespace, and dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#metric CloudwatchAlarm#metric}
	Metric *CloudwatchAlarmMetricsMetricStatMetric `field:"required" json:"metric" yaml:"metric"`
	// The granularity, in seconds, of the returned data points.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#period CloudwatchAlarm#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The statistic to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#stat CloudwatchAlarm#stat}
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#unit CloudwatchAlarm#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

