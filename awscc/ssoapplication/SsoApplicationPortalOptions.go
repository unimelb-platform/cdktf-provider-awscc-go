package ssoapplication


type SsoApplicationPortalOptions struct {
	// A structure that describes the sign-in options for the access portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#sign_in_options SsoApplication#sign_in_options}
	SignInOptions *SsoApplicationPortalOptionsSignInOptions `field:"optional" json:"signInOptions" yaml:"signInOptions"`
	// Indicates whether this application is visible in the access portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application#visibility SsoApplication#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

