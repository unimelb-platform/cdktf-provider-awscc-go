package iotwirelessdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotwirelessDestinationConfig struct {
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
	// Destination expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_destination#expression IotwirelessDestination#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Must be RuleName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_destination#expression_type IotwirelessDestination#expression_type}
	ExpressionType *string `field:"required" json:"expressionType" yaml:"expressionType"`
	// Unique name of destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_destination#name IotwirelessDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Destination description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_destination#description IotwirelessDestination#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS role ARN that grants access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_destination#role_arn IotwirelessDestination#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A list of key-value pairs that contain metadata for the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_destination#tags IotwirelessDestination#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

