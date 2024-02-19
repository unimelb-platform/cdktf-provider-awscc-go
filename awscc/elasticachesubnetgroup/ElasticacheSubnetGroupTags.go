package elasticachesubnetgroup


type ElasticacheSubnetGroupTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_subnet_group#key ElasticacheSubnetGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_subnet_group#value ElasticacheSubnetGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

