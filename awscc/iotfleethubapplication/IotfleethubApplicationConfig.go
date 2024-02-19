package iotfleethubapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotfleethubApplicationConfig struct {
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
	// Application Name, should be between 1 and 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleethub_application#application_name IotfleethubApplication#application_name}
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The ARN of the role that the web application assumes when it interacts with AWS IoT Core.
	//
	// For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleethub_application#role_arn IotfleethubApplication#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Application Description, should be between 1 and 2048 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleethub_application#application_description IotfleethubApplication#application_description}
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// A list of key-value pairs that contain metadata for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotfleethub_application#tags IotfleethubApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

