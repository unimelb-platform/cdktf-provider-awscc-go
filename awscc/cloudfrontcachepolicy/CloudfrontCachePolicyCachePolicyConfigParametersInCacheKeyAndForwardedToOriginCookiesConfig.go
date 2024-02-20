package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#cookie_behavior CloudfrontCachePolicy#cookie_behavior}.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#cookies CloudfrontCachePolicy#cookies}.
	Cookies *[]*string `field:"optional" json:"cookies" yaml:"cookies"`
}

