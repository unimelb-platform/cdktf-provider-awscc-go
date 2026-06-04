package sagemakerpartnerapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerPartnerAppConfig struct {
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
	// The Auth type of PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#auth_type SagemakerPartnerApp#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The execution role for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#execution_role_arn SagemakerPartnerApp#execution_role_arn}
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// A name for the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#name SagemakerPartnerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The tier of the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#tier SagemakerPartnerApp#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The type of PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#type SagemakerPartnerApp#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A collection of settings that specify the maintenance schedule for the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#application_config SagemakerPartnerApp#application_config}
	ApplicationConfig *SagemakerPartnerAppApplicationConfig `field:"optional" json:"applicationConfig" yaml:"applicationConfig"`
	// The client token for the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#client_token SagemakerPartnerApp#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Enables IAM Session based Identity for PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#enable_iam_session_based_identity SagemakerPartnerApp#enable_iam_session_based_identity}
	EnableIamSessionBasedIdentity interface{} `field:"optional" json:"enableIamSessionBasedIdentity" yaml:"enableIamSessionBasedIdentity"`
	// The AWS KMS customer managed key used to encrypt the data associated with the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#kms_key_id SagemakerPartnerApp#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A collection of settings that specify the maintenance schedule for the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#maintenance_config SagemakerPartnerApp#maintenance_config}
	MaintenanceConfig *SagemakerPartnerAppMaintenanceConfig `field:"optional" json:"maintenanceConfig" yaml:"maintenanceConfig"`
	// A list of tags to apply to the PartnerApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_partner_app#tags SagemakerPartnerApp#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

