package logsmetricfilter


type LogsMetricFilterMetricTransformations struct {
	// The name of the CloudWatch metric. Metric name must be in ASCII format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#metric_name LogsMetricFilter#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The namespace of the CloudWatch metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#metric_namespace LogsMetricFilter#metric_namespace}
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The value to publish to the CloudWatch metric when a filter pattern matches a log event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#metric_value LogsMetricFilter#metric_value}
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// The value to emit when a filter pattern does not match a log event. This value can be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#default_value LogsMetricFilter#default_value}
	DefaultValue *float64 `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Dimensions are the key-value pairs that further define a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#dimensions LogsMetricFilter#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The unit to assign to the metric. If you omit this, the unit is set as None.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#unit LogsMetricFilter#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

