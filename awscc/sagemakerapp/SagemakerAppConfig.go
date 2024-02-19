package sagemakerapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerAppConfig struct {
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
	// The name of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app#app_name SagemakerApp#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The type of app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app#app_type SagemakerApp#app_type}
	AppType *string `field:"required" json:"appType" yaml:"appType"`
	// The domain ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app#domain_id SagemakerApp#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The user profile name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app#user_profile_name SagemakerApp#user_profile_name}
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app#resource_spec SagemakerApp#resource_spec}
	ResourceSpec *SagemakerAppResourceSpec `field:"optional" json:"resourceSpec" yaml:"resourceSpec"`
	// A list of tags to apply to the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app#tags SagemakerApp#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

