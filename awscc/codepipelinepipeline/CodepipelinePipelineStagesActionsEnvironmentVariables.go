package codepipelinepipeline


type CodepipelinePipelineStagesActionsEnvironmentVariables struct {
	// The name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#name CodepipelinePipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#type CodepipelinePipeline#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#value CodepipelinePipeline#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

