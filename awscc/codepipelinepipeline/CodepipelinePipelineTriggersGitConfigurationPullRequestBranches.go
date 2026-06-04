package codepipelinepipeline


type CodepipelinePipelineTriggersGitConfigurationPullRequestBranches struct {
	// The list of patterns of Git branches that, when a commit is pushed, are to be excluded from starting the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#excludes CodepipelinePipeline#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// The list of patterns of Git branches that, when a commit is pushed, are to be included as criteria that starts the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codepipeline_pipeline#includes CodepipelinePipeline#includes}
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}

