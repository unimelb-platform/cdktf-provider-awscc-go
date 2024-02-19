package logsdelivery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsDeliveryConfig struct {
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
	// The ARN of the delivery destination that is associated with this delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery#delivery_destination_arn LogsDelivery#delivery_destination_arn}
	DeliveryDestinationArn *string `field:"required" json:"deliveryDestinationArn" yaml:"deliveryDestinationArn"`
	// The name of the delivery source that is associated with this delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery#delivery_source_name LogsDelivery#delivery_source_name}
	DeliverySourceName *string `field:"required" json:"deliverySourceName" yaml:"deliverySourceName"`
	// The tags that have been assigned to this delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery#tags LogsDelivery#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

