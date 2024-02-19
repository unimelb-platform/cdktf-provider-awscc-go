package cloudfrontdistribution


type CloudfrontDistributionTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#key CloudfrontDistribution#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#value CloudfrontDistribution#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

