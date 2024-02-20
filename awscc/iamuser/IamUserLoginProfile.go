package iamuser


type IamUserLoginProfile struct {
	// The user's password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#password IamUser#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies whether the user is required to set a new password on next sign-in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iam_user#password_reset_required IamUser#password_reset_required}
	PasswordResetRequired interface{} `field:"optional" json:"passwordResetRequired" yaml:"passwordResetRequired"`
}

