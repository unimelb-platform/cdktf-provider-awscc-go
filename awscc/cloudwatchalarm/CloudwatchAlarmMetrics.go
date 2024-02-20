package cloudwatchalarm


type CloudwatchAlarmMetrics struct {
	// A short name used to tie this object to the results in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#id CloudwatchAlarm#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#account_id CloudwatchAlarm#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The math expression to be performed on the returned data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#expression CloudwatchAlarm#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A human-readable label for this metric or expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#label CloudwatchAlarm#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The metric to be returned, along with statistics, period, and units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#metric_stat CloudwatchAlarm#metric_stat}
	MetricStat *CloudwatchAlarmMetricsMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// The period in seconds, over which the statistic is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#period CloudwatchAlarm#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// This option indicates whether to return the timestamps and raw data values of this metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#return_data CloudwatchAlarm#return_data}
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

