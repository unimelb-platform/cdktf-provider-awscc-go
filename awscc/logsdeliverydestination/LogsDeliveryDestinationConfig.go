package logsdeliverydestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsDeliveryDestinationConfig struct {
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
	// The name of this delivery destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_delivery_destination#name LogsDeliveryDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 51200
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_delivery_destination#delivery_destination_policy LogsDeliveryDestination#delivery_destination_policy}
	DeliveryDestinationPolicy *LogsDeliveryDestinationDeliveryDestinationPolicy `field:"optional" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
	// The ARN of the Amazon Web Services destination that this delivery destination represents.
	//
	// That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_delivery_destination#destination_resource_arn LogsDeliveryDestination#destination_resource_arn}
	DestinationResourceArn *string `field:"optional" json:"destinationResourceArn" yaml:"destinationResourceArn"`
	// The format of the logs that are sent to this delivery destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_delivery_destination#output_format LogsDeliveryDestination#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The tags that have been assigned to this delivery destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_delivery_destination#tags LogsDeliveryDestination#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

