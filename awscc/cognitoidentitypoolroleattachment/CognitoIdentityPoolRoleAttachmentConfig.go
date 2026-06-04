package cognitoidentitypoolroleattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoIdentityPoolRoleAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#identity_pool_id CognitoIdentityPoolRoleAttachment#identity_pool_id}.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#role_mappings CognitoIdentityPoolRoleAttachment#role_mappings}.
	RoleMappings interface{} `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#roles CognitoIdentityPoolRoleAttachment#roles}.
	Roles *map[string]*string `field:"optional" json:"roles" yaml:"roles"`
}

