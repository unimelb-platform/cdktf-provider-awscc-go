package elasticacheuser


type ElasticacheUserAuthenticationMode struct {
	// Authentication Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_user#type ElasticacheUser#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Passwords used for this user account. You can create up to two passwords for each user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_user#passwords ElasticacheUser#passwords}
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

