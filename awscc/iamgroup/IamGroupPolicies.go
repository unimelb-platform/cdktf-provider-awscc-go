package iamgroup


type IamGroupPolicies struct {
	// The policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_group#policy_document IamGroup#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iam_group#policy_name IamGroup#policy_name}
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

