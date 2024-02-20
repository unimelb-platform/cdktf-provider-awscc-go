package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#query_string_behavior CloudfrontCachePolicy#query_string_behavior}.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#query_strings CloudfrontCachePolicy#query_strings}.
	QueryStrings *[]*string `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

