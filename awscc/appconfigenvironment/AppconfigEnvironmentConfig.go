package appconfigenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppconfigEnvironmentConfig struct {
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
	// The application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#application_id AppconfigEnvironment#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A name for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#name AppconfigEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#description AppconfigEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Amazon CloudWatch alarms to monitor during the deployment process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#monitors AppconfigEnvironment#monitors}
	Monitors interface{} `field:"optional" json:"monitors" yaml:"monitors"`
	// Metadata to assign to the environment.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#tags AppconfigEnvironment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

