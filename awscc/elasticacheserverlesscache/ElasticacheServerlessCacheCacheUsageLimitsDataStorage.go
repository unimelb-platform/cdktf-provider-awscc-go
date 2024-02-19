package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimitsDataStorage struct {
	// The maximum cached data capacity of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_serverless_cache#maximum ElasticacheServerlessCache#maximum}
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// The unix of cached data capacity of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_serverless_cache#unit ElasticacheServerlessCache#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

