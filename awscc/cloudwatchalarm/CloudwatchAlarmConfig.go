package cloudwatchalarm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchAlarmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#comparison_operator CloudwatchAlarm#comparison_operator}
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of periods over which data is compared to the specified threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#evaluation_periods CloudwatchAlarm#evaluation_periods}
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#actions_enabled CloudwatchAlarm#actions_enabled}
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#alarm_actions CloudwatchAlarm#alarm_actions}
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#alarm_description CloudwatchAlarm#alarm_description}
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#alarm_name CloudwatchAlarm#alarm_name}
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#datapoints_to_alarm CloudwatchAlarm#datapoints_to_alarm}
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The dimensions for the metric associated with the alarm.
	//
	// For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#dimensions CloudwatchAlarm#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Used only for alarms based on percentiles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#evaluate_low_sample_count_percentile CloudwatchAlarm#evaluate_low_sample_count_percentile}
	EvaluateLowSampleCountPercentile *string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#extended_statistic CloudwatchAlarm#extended_statistic}
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#insufficient_data_actions CloudwatchAlarm#insufficient_data_actions}
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The name of the metric associated with the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#metric_name CloudwatchAlarm#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// An array that enables you to create an alarm based on the result of a metric math expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#metrics CloudwatchAlarm#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The namespace of the metric associated with the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#namespace CloudwatchAlarm#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The actions to execute when this alarm transitions to the OK state from any other state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#ok_actions CloudwatchAlarm#ok_actions}
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// The period in seconds, over which the statistic is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#period CloudwatchAlarm#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic for the metric associated with the alarm, other than percentile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#statistic CloudwatchAlarm#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#threshold CloudwatchAlarm#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#threshold_metric_id CloudwatchAlarm#threshold_metric_id}
	ThresholdMetricId *string `field:"optional" json:"thresholdMetricId" yaml:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#treat_missing_data CloudwatchAlarm#treat_missing_data}
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
	// The unit of the metric associated with the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_alarm#unit CloudwatchAlarm#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

