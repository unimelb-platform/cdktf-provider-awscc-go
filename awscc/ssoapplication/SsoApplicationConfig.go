package ssoapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoApplicationConfig struct {
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
	// The ARN of the application provider under which the operation will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#application_provider_arn SsoApplication#application_provider_arn}
	ApplicationProviderArn *string `field:"required" json:"applicationProviderArn" yaml:"applicationProviderArn"`
	// The ARN of the instance of IAM Identity Center under which the operation will run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#instance_arn SsoApplication#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name you want to assign to this Identity Center (SSO) Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#name SsoApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description information for the Identity Center (SSO) Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#description SsoApplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that describes the options for the portal associated with an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#portal_options SsoApplication#portal_options}
	PortalOptions *SsoApplicationPortalOptions `field:"optional" json:"portalOptions" yaml:"portalOptions"`
	// Specifies whether the application is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#status SsoApplication#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#tags SsoApplication#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

