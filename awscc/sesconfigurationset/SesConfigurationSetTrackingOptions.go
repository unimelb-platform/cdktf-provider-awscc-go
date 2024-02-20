package sesconfigurationset


type SesConfigurationSetTrackingOptions struct {
	// The domain to use for tracking open and click events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#custom_redirect_domain SesConfigurationSet#custom_redirect_domain}
	CustomRedirectDomain *string `field:"optional" json:"customRedirectDomain" yaml:"customRedirectDomain"`
}

