package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigRestrictionsGeoRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#restriction_type CloudfrontDistribution#restriction_type}.
	RestrictionType *string `field:"required" json:"restrictionType" yaml:"restrictionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#locations CloudfrontDistribution#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
}

