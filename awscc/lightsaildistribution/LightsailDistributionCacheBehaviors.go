package lightsaildistribution


type LightsailDistributionCacheBehaviors struct {
	// The cache behavior for the specified path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#behavior LightsailDistribution#behavior}
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// The path to a directory or file to cached, or not cache.
	//
	// Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_distribution#path LightsailDistribution#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

