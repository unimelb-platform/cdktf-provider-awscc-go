package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#geo_restriction CloudfrontDistribution#geo_restriction}.
	GeoRestriction *CloudfrontDistributionDistributionConfigRestrictionsGeoRestriction `field:"required" json:"geoRestriction" yaml:"geoRestriction"`
}

