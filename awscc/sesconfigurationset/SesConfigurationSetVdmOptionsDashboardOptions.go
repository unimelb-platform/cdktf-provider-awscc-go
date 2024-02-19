package sesconfigurationset


type SesConfigurationSetVdmOptionsDashboardOptions struct {
	// Whether emails sent with this configuration set have engagement tracking enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set#engagement_metrics SesConfigurationSet#engagement_metrics}
	EngagementMetrics *string `field:"required" json:"engagementMetrics" yaml:"engagementMetrics"`
}

