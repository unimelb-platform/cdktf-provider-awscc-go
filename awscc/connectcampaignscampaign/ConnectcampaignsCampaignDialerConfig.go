package connectcampaignscampaign


type ConnectcampaignsCampaignDialerConfig struct {
	// Agentless Dialer config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connectcampaigns_campaign#agentless_dialer_config ConnectcampaignsCampaign#agentless_dialer_config}
	AgentlessDialerConfig *ConnectcampaignsCampaignDialerConfigAgentlessDialerConfig `field:"optional" json:"agentlessDialerConfig" yaml:"agentlessDialerConfig"`
	// Predictive Dialer config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connectcampaigns_campaign#predictive_dialer_config ConnectcampaignsCampaign#predictive_dialer_config}
	PredictiveDialerConfig *ConnectcampaignsCampaignDialerConfigPredictiveDialerConfig `field:"optional" json:"predictiveDialerConfig" yaml:"predictiveDialerConfig"`
	// Progressive Dialer config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connectcampaigns_campaign#progressive_dialer_config ConnectcampaignsCampaign#progressive_dialer_config}
	ProgressiveDialerConfig *ConnectcampaignsCampaignDialerConfigProgressiveDialerConfig `field:"optional" json:"progressiveDialerConfig" yaml:"progressiveDialerConfig"`
}

