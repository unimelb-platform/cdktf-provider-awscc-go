package sesconfigurationset


type SesConfigurationSetSuppressionOptions struct {
	// A list that contains the reasons that email addresses are automatically added to the suppression list for your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#suppressed_reasons SesConfigurationSet#suppressed_reasons}
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
}

