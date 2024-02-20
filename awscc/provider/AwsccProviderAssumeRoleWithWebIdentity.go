package provider


type AwsccProviderAssumeRoleWithWebIdentity struct {
	// Amazon Resource Name (ARN) of the IAM Role to assume. Can also be set with the environment variable `AWS_ROLE_ARN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#role_arn AwsccProvider#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Duration of the assume role session.
	//
	// You can provide a value from 15 minutes up to the maximum session duration setting for the role. A sequence of numbers with a unit suffix, "h" for hour, "m" for minute, and "s" for second. Default value is 1h0m0s
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#duration AwsccProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// IAM policy in JSON format to use as a session policy.
	//
	// The effective permissions for the session will be the intersection between this polcy and the role's policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#policy AwsccProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Amazon Resource Names (ARNs) of IAM Policies to use as managed session policies.
	//
	// The effective permissions for the session will be the intersection between these polcy and the role's policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#policy_arns AwsccProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Session name to use when assuming the role. Can also be set with the environment variable `AWS_ROLE_SESSION_NAME`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#session_name AwsccProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// The value of a web identity token from an OpenID Connect (OIDC) or OAuth provider.
	//
	// One of `web_identity_token` or `web_identity_token_file` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#web_identity_token AwsccProvider#web_identity_token}
	WebIdentityToken *string `field:"optional" json:"webIdentityToken" yaml:"webIdentityToken"`
	// File containing a web identity token from an OpenID Connect (OIDC) or OAuth provider.
	//
	// Can also be set with the  environment variable`AWS_WEB_IDENTITY_TOKEN_FILE`. One of `web_identity_token_file` or `web_identity_token` is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#web_identity_token_file AwsccProvider#web_identity_token_file}
	WebIdentityTokenFile *string `field:"optional" json:"webIdentityTokenFile" yaml:"webIdentityTokenFile"`
}

