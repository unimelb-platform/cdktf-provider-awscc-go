package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#header_behavior CloudfrontCachePolicy#header_behavior}.
	HeaderBehavior *string `field:"required" json:"headerBehavior" yaml:"headerBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#headers CloudfrontCachePolicy#headers}.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
}

