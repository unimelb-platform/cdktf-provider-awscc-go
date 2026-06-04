package codepipelinepipeline


type CodepipelinePipelineArtifactStores struct {
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#artifact_store CodepipelinePipeline#artifact_store}
	ArtifactStore *CodepipelinePipelineArtifactStoresArtifactStore `field:"optional" json:"artifactStore" yaml:"artifactStore"`
	// The action declaration's AWS Region, such as us-east-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#region CodepipelinePipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

