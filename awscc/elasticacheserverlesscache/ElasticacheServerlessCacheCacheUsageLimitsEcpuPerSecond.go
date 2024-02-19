package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond struct {
	// The maximum ECPU per second of the Serverless Cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_serverless_cache#maximum ElasticacheServerlessCache#maximum}
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
}

