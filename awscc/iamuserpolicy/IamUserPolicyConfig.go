package iamuserpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamUserPolicyConfig struct {
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
	// The name of the policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user_policy#policy_name IamUserPolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the user to associate the policy with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user_policy#user_name IamUserPolicy#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// The policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user_policy#policy_document IamUserPolicy#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

