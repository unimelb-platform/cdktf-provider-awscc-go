package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfig struct {
	// Email Channel Subtype config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#email Connectcampaignsv2Campaign#email}
	Email *Connectcampaignsv2CampaignChannelSubtypeConfigEmail `field:"optional" json:"email" yaml:"email"`
	// SMS Channel Subtype config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#sms Connectcampaignsv2Campaign#sms}
	Sms *Connectcampaignsv2CampaignChannelSubtypeConfigSms `field:"optional" json:"sms" yaml:"sms"`
	// Telephony Channel Subtype config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#telephony Connectcampaignsv2Campaign#telephony}
	Telephony *Connectcampaignsv2CampaignChannelSubtypeConfigTelephony `field:"optional" json:"telephony" yaml:"telephony"`
}

