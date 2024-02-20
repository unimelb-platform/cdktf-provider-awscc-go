package iamrole


type IamRolePolicies struct {
	// The policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#policy_document IamRole#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_role#policy_name IamRole#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

