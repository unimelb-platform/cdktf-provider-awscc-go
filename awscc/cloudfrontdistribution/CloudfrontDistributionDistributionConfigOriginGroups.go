package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigOriginGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#quantity CloudfrontDistribution#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#items CloudfrontDistribution#items}.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

