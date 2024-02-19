package cloudwatchmetricstream


type CloudwatchMetricStreamTags struct {
	// A unique identifier for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#key CloudwatchMetricStream#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// String which you can use to describe or define the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudwatch_metric_stream#value CloudwatchMetricStream#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

