package logsdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsDestinationConfig struct {
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
	// The name of the destination resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_destination#destination_name LogsDestination#destination_name}
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_destination#role_arn LogsDestination#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the physical target where the log events are delivered (for example, a Kinesis stream).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_destination#target_arn LogsDestination#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_destination#destination_policy LogsDestination#destination_policy}
	DestinationPolicy *string `field:"optional" json:"destinationPolicy" yaml:"destinationPolicy"`
}

