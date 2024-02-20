package lightsaildistribution


type LightsailDistributionDefaultCacheBehavior struct {
	// The cache behavior of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#behavior LightsailDistribution#behavior}
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
}

