package lightsaildistribution


type LightsailDistributionCacheBehaviorSettingsForwardedQueryStrings struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#option LightsailDistribution#option}
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// The specific query strings that the distribution forwards to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#query_strings_allow_list LightsailDistribution#query_strings_allow_list}
	QueryStringsAllowList *[]*string `field:"optional" json:"queryStringsAllowList" yaml:"queryStringsAllowList"`
}

