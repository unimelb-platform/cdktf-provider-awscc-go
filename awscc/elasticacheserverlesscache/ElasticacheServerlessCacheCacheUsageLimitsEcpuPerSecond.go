package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond struct {
	// The maximum ECPU per second of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticache_serverless_cache#maximum ElasticacheServerlessCache#maximum}
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The minimum ECPU per second of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticache_serverless_cache#minimum ElasticacheServerlessCache#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

