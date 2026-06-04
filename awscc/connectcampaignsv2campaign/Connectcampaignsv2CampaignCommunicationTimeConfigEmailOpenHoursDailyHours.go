package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigEmailOpenHoursDailyHours struct {
	// Day of week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#key Connectcampaignsv2Campaign#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// List of time range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#value Connectcampaignsv2Campaign#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

