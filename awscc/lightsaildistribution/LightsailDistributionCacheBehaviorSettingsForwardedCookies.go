package lightsaildistribution


type LightsailDistributionCacheBehaviorSettingsForwardedCookies struct {
	// The specific cookies to forward to your distribution's origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#cookies_allow_list LightsailDistribution#cookies_allow_list}
	CookiesAllowList *[]*string `field:"optional" json:"cookiesAllowList" yaml:"cookiesAllowList"`
	// Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#option LightsailDistribution#option}
	Option *string `field:"optional" json:"option" yaml:"option"`
}

