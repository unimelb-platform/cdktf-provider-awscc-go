package codepipelinepipeline


type CodepipelinePipelineStagesOnFailureRetryConfiguration struct {
	// The specified retry mode type for the given stage.
	//
	// FAILED_ACTIONS will retry only the failed actions. ALL_ACTIONS will retry both failed and successful
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#retry_mode CodepipelinePipeline#retry_mode}
	RetryMode *string `field:"optional" json:"retryMode" yaml:"retryMode"`
}

