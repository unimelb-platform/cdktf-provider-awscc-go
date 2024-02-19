package cloudfrontcachepolicy


type CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOrigin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#cookies_config CloudfrontCachePolicy#cookies_config}.
	CookiesConfig *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginCookiesConfig `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_gzip CloudfrontCachePolicy#enable_accept_encoding_gzip}.
	EnableAcceptEncodingGzip interface{} `field:"required" json:"enableAcceptEncodingGzip" yaml:"enableAcceptEncodingGzip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#headers_config CloudfrontCachePolicy#headers_config}.
	HeadersConfig *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginHeadersConfig `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#query_strings_config CloudfrontCachePolicy#query_strings_config}.
	QueryStringsConfig *CloudfrontCachePolicyCachePolicyConfigParametersInCacheKeyAndForwardedToOriginQueryStringsConfig `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cache_policy#enable_accept_encoding_brotli CloudfrontCachePolicy#enable_accept_encoding_brotli}.
	EnableAcceptEncodingBrotli interface{} `field:"optional" json:"enableAcceptEncodingBrotli" yaml:"enableAcceptEncodingBrotli"`
}

