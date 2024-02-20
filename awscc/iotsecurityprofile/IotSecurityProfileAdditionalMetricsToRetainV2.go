package iotsecurityprofile


type IotSecurityProfileAdditionalMetricsToRetainV2 struct {
	// What is measured by the behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#metric IotSecurityProfile#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// Flag to enable/disable metrics export for metric to be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#export_metric IotSecurityProfile#export_metric}
	ExportMetric interface{} `field:"optional" json:"exportMetric" yaml:"exportMetric"`
	// The dimension of a metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#metric_dimension IotSecurityProfile#metric_dimension}
	MetricDimension *IotSecurityProfileAdditionalMetricsToRetainV2MetricDimension `field:"optional" json:"metricDimension" yaml:"metricDimension"`
}

