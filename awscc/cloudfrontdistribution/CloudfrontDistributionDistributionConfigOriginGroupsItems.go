package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigOriginGroupsItems struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#failover_criteria CloudfrontDistribution#failover_criteria}.
	FailoverCriteria *CloudfrontDistributionDistributionConfigOriginGroupsItemsFailoverCriteria `field:"required" json:"failoverCriteria" yaml:"failoverCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#id CloudfrontDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#members CloudfrontDistribution#members}.
	Members *CloudfrontDistributionDistributionConfigOriginGroupsItemsMembers `field:"required" json:"members" yaml:"members"`
}

