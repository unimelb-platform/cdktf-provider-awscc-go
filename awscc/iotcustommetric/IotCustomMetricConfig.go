package iotcustommetric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotCustomMetricConfig struct {
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
	// The type of the custom metric. Types include string-list, ip-address-list, number-list, and number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_custom_metric#metric_type IotCustomMetric#metric_type}
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// Field represents a friendly name in the console for the custom metric;
	//
	// it doesn't have to be unique. Don't use this name as the metric identifier in the device metric report. Can be updated once defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_custom_metric#display_name IotCustomMetric#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The name of the custom metric.
	//
	// This will be used in the metric report submitted from the device/thing. Shouldn't begin with aws: . Cannot be updated once defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_custom_metric#metric_name IotCustomMetric#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_custom_metric#tags IotCustomMetric#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

