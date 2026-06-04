package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfig struct {
	// Time Zone Id in the IANA format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#default_time_zone Connectcampaignsv2Campaign#default_time_zone}
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// Local TimeZone Detection method list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#local_time_zone_detection Connectcampaignsv2Campaign#local_time_zone_detection}
	LocalTimeZoneDetection *[]*string `field:"optional" json:"localTimeZoneDetection" yaml:"localTimeZoneDetection"`
}

