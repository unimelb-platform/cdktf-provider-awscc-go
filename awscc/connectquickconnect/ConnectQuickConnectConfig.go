package connectquickconnect

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectQuickConnectConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#instance_arn ConnectQuickConnect#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the quick connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#name ConnectQuickConnect#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configuration settings for the quick connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#quick_connect_config ConnectQuickConnect#quick_connect_config}
	QuickConnectConfig *ConnectQuickConnectQuickConnectConfig `field:"required" json:"quickConnectConfig" yaml:"quickConnectConfig"`
	// The description of the quick connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#description ConnectQuickConnect#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#tags ConnectQuickConnect#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

