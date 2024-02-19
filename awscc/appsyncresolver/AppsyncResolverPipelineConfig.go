package appsyncresolver


type AppsyncResolverPipelineConfig struct {
	// A list of Function objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#functions AppsyncResolver#functions}
	Functions *[]*string `field:"optional" json:"functions" yaml:"functions"`
}

