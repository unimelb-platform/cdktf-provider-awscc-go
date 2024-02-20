package iamuser


type IamUserPolicies struct {
	// The policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#policy_document IamUser#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#policy_name IamUser#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

