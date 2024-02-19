package cognitoidentitypoolprincipaltag

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoIdentityPoolPrincipalTagConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool_principal_tag#identity_pool_id CognitoIdentityPoolPrincipalTag#identity_pool_id}.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool_principal_tag#identity_provider_name CognitoIdentityPoolPrincipalTag#identity_provider_name}.
	IdentityProviderName *string `field:"required" json:"identityProviderName" yaml:"identityProviderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool_principal_tag#principal_tags CognitoIdentityPoolPrincipalTag#principal_tags}.
	PrincipalTags *string `field:"optional" json:"principalTags" yaml:"principalTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool_principal_tag#use_defaults CognitoIdentityPoolPrincipalTag#use_defaults}.
	UseDefaults interface{} `field:"optional" json:"useDefaults" yaml:"useDefaults"`
}

