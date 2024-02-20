package sagemakeruserprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileConfig struct {
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
	// The ID of the associated Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#domain_id SagemakerUserProfile#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// A name for the UserProfile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#user_profile_name SagemakerUserProfile#user_profile_name}
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// A specifier for the type of value specified in SingleSignOnUserValue.
	//
	// Currently, the only supported value is "UserName". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#single_sign_on_user_identifier SagemakerUserProfile#single_sign_on_user_identifier}
	SingleSignOnUserIdentifier *string `field:"optional" json:"singleSignOnUserIdentifier" yaml:"singleSignOnUserIdentifier"`
	// The username of the associated AWS Single Sign-On User for this UserProfile.
	//
	// If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#single_sign_on_user_value SagemakerUserProfile#single_sign_on_user_value}
	SingleSignOnUserValue *string `field:"optional" json:"singleSignOnUserValue" yaml:"singleSignOnUserValue"`
	// A list of tags to apply to the user profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#tags SagemakerUserProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A collection of settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#user_settings SagemakerUserProfile#user_settings}
	UserSettings *SagemakerUserProfileUserSettings `field:"optional" json:"userSettings" yaml:"userSettings"`
}

