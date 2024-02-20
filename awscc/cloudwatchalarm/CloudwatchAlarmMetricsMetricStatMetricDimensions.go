package cloudwatchalarm


type CloudwatchAlarmMetricsMetricStatMetricDimensions struct {
	// The name of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#name CloudwatchAlarm#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value for the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#value CloudwatchAlarm#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

