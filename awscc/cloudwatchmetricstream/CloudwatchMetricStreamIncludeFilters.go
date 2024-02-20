package cloudwatchmetricstream


type CloudwatchMetricStreamIncludeFilters struct {
	// Only metrics with Namespace matching this value will be streamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#namespace CloudwatchMetricStream#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Only metrics with MetricNames matching these values will be streamed. Must be set together with Namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#metric_names CloudwatchMetricStream#metric_names}
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

