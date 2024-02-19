package cloudwatchalarm


type CloudwatchAlarmMetricsMetricStatMetric struct {
	// The dimensions for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#dimensions CloudwatchAlarm#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#metric_name CloudwatchAlarm#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#namespace CloudwatchAlarm#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

