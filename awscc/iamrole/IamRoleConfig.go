package iamrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamRoleConfig struct {
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
	// The trust policy that is associated with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#assume_role_policy_document IamRole#assume_role_policy_document}
	AssumeRolePolicyDocument *string `field:"required" json:"assumeRolePolicyDocument" yaml:"assumeRolePolicyDocument"`
	// A description of the role that you provide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#description IamRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#managed_policy_arns IamRole#managed_policy_arns}
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// The maximum session duration (in seconds) that you want to set for the specified role.
	//
	// If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#max_session_duration IamRole#max_session_duration}
	MaxSessionDuration *float64 `field:"optional" json:"maxSessionDuration" yaml:"maxSessionDuration"`
	// The path to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#path IamRole#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The ARN of the policy used to set the permissions boundary for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#permissions_boundary IamRole#permissions_boundary}
	PermissionsBoundary *string `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Adds or updates an inline policy document that is embedded in the specified IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#policies IamRole#policies}
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// A name for the IAM role, up to 64 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#role_name IamRole#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// A list of tags that are attached to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#tags IamRole#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

