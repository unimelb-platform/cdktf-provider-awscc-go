package sesemailidentity


type SesEmailIdentityFeedbackAttributes struct {
	// If the value is true, you receive email notifications when bounce or complaint events occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity#email_forwarding_enabled SesEmailIdentity#email_forwarding_enabled}
	EmailForwardingEnabled interface{} `field:"optional" json:"emailForwardingEnabled" yaml:"emailForwardingEnabled"`
}

