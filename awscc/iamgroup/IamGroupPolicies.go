package iamgroup


type IamGroupPolicies struct {
	// The policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_group#policy_document IamGroup#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_group#policy_name IamGroup#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

