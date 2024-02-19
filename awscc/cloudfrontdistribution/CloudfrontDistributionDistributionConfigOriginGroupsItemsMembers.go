package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigOriginGroupsItemsMembers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#items CloudfrontDistribution#items}.
	Items interface{} `field:"required" json:"items" yaml:"items"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#quantity CloudfrontDistribution#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}

