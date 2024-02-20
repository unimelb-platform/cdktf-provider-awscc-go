package iamuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamUserConfig struct {
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
	// A list of group names to which you want to add the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#groups IamUser#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Creates a password for the specified IAM user.
	//
	// A password allows an IAM user to access AWS services through the AWS Management Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#login_profile IamUser#login_profile}
	LoginProfile *IamUserLoginProfile `field:"optional" json:"loginProfile" yaml:"loginProfile"`
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#managed_policy_arns IamUser#managed_policy_arns}
	ManagedPolicyArns *[]*string `field:"optional" json:"managedPolicyArns" yaml:"managedPolicyArns"`
	// The path to the user.
	//
	// For more information about paths, see IAM identifiers in the IAM User Guide. The ARN of the policy used to set the permissions boundary for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#path IamUser#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#permissions_boundary IamUser#permissions_boundary}
	PermissionsBoundary *string `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Adds or updates an inline policy document that is embedded in the specified IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#policies IamUser#policies}
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// A list of tags that are associated with the user.
	//
	// For more information about tagging, see Tagging IAM resources in the IAM User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#tags IamUser#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The friendly name identifying the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#user_name IamUser#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

