package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundMode struct {
	// Agentless config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#agentless_config Connectcampaignsv2Campaign#agentless_config}
	AgentlessConfig *string `field:"optional" json:"agentlessConfig" yaml:"agentlessConfig"`
	// Predictive config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#predictive_config Connectcampaignsv2Campaign#predictive_config}
	PredictiveConfig *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePredictiveConfig `field:"optional" json:"predictiveConfig" yaml:"predictiveConfig"`
	// Progressive config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#progressive_config Connectcampaignsv2Campaign#progressive_config}
	ProgressiveConfig *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModeProgressiveConfig `field:"optional" json:"progressiveConfig" yaml:"progressiveConfig"`
}

