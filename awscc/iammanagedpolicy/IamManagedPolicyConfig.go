package iammanagedpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamManagedPolicyConfig struct {
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
	// The JSON policy document that you want to use as the content for the new policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#policy_document IamManagedPolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// A friendly description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#description IamManagedPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name (friendly name, not ARN) of the group to attach the policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#groups IamManagedPolicy#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// The friendly name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#managed_policy_name IamManagedPolicy#managed_policy_name}
	ManagedPolicyName *string `field:"optional" json:"managedPolicyName" yaml:"managedPolicyName"`
	// The path for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#path IamManagedPolicy#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The name (friendly name, not ARN) of the role to attach the policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#roles IamManagedPolicy#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// The name (friendly name, not ARN) of the IAM user to attach the policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_managed_policy#users IamManagedPolicy#users}
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

