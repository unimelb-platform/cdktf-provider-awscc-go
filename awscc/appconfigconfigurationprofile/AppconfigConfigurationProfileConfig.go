package appconfigconfigurationprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppconfigConfigurationProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#application_id AppconfigConfigurationProfile#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A URI to locate the configuration.
	//
	// You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#location_uri AppconfigConfigurationProfile#location_uri}
	LocationUri *string `field:"required" json:"locationUri" yaml:"locationUri"`
	// A name for the configuration profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#name AppconfigConfigurationProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the configuration profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#description AppconfigConfigurationProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#kms_key_identifier AppconfigConfigurationProfile#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The ARN of an IAM role with permission to access the configuration at the specified LocationUri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#retrieval_role_arn AppconfigConfigurationProfile#retrieval_role_arn}
	RetrievalRoleArn *string `field:"optional" json:"retrievalRoleArn" yaml:"retrievalRoleArn"`
	// Metadata to assign to the configuration profile.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#tags AppconfigConfigurationProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of configurations contained in the profile.
	//
	// When calling this API, enter one of the following values for Type: AWS.AppConfig.FeatureFlags, AWS.Freeform
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#type AppconfigConfigurationProfile#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// A list of methods for validating the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#validators AppconfigConfigurationProfile#validators}
	Validators interface{} `field:"optional" json:"validators" yaml:"validators"`
}

