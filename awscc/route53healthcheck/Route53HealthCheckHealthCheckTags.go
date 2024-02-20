package route53healthcheck


type Route53HealthCheckHealthCheckTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_health_check#key Route53HealthCheck#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_health_check#value Route53HealthCheck#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

