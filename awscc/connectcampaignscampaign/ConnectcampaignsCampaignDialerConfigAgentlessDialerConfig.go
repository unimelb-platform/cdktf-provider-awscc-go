package connectcampaignscampaign


type ConnectcampaignsCampaignDialerConfigAgentlessDialerConfig struct {
	// Allocates dialing capacity for this campaign between multiple active campaigns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connectcampaigns_campaign#dialing_capacity ConnectcampaignsCampaign#dialing_capacity}
	DialingCapacity *float64 `field:"optional" json:"dialingCapacity" yaml:"dialingCapacity"`
}

