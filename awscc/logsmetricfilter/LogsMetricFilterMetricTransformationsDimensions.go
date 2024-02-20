package logsmetricfilter


type LogsMetricFilterMetricTransformationsDimensions struct {
	// The key of the dimension. Maximum length of 255.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#key LogsMetricFilter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the dimension. Maximum length of 255.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_metric_filter#value LogsMetricFilter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

