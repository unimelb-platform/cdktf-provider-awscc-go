package appsyncresolver


type AppsyncResolverCachingConfig struct {
	// The TTL in seconds for a resolver that has caching activated. Valid values are 1-36.00 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#ttl AppsyncResolver#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// The caching keys for a resolver that has caching activated.
	//
	// Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#caching_keys AppsyncResolver#caching_keys}
	CachingKeys *[]*string `field:"optional" json:"cachingKeys" yaml:"cachingKeys"`
}

