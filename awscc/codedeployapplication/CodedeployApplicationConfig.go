package codedeployapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodedeployApplicationConfig struct {
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
	// A name for the application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codedeploy_application#application_name CodedeployApplication#application_name}
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// The compute platform that CodeDeploy deploys the application to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codedeploy_application#compute_platform CodedeployApplication#compute_platform}
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// The metadata that you apply to CodeDeploy applications to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codedeploy_application#tags CodedeployApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

