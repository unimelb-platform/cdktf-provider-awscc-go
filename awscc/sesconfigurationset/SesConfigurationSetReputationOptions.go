package sesconfigurationset


type SesConfigurationSetReputationOptions struct {
	// If true , tracking of reputation metrics is enabled for the configuration set.
	//
	// If false , tracking of reputation metrics is disabled for the configuration set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#reputation_metrics_enabled SesConfigurationSet#reputation_metrics_enabled}
	ReputationMetricsEnabled interface{} `field:"optional" json:"reputationMetricsEnabled" yaml:"reputationMetricsEnabled"`
}

