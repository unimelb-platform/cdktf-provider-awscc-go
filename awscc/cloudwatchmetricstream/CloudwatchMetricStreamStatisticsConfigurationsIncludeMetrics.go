package cloudwatchmetricstream


type CloudwatchMetricStreamStatisticsConfigurationsIncludeMetrics struct {
	// The name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudwatch_metric_stream#metric_name CloudwatchMetricStream#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudwatch_metric_stream#namespace CloudwatchMetricStream#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

