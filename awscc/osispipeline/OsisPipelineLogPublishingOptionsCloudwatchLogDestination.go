package osispipeline


type OsisPipelineLogPublishingOptionsCloudwatchLogDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/osis_pipeline#log_group OsisPipeline#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

