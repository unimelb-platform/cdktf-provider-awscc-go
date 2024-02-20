package connectcampaignscampaign


type ConnectcampaignsCampaignTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connectcampaigns_campaign#key ConnectcampaignsCampaign#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that's 1 to 256 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connectcampaigns_campaign#value ConnectcampaignsCampaign#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

