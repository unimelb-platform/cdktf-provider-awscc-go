package ec2networkperformancemetricsubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkPerformanceMetricSubscriptionConfig struct {
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
	// The target Region or Availability Zone for the metric to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_performance_metric_subscription#destination Ec2NetworkPerformanceMetricSubscription#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The metric type to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_performance_metric_subscription#metric Ec2NetworkPerformanceMetricSubscription#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The starting Region or Availability Zone for metric to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_performance_metric_subscription#source Ec2NetworkPerformanceMetricSubscription#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The statistic to subscribe to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_performance_metric_subscription#statistic Ec2NetworkPerformanceMetricSubscription#statistic}
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

