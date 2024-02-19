package sagemakerpipeline


type SagemakerPipelineTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_pipeline#key SagemakerPipeline#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_pipeline#value SagemakerPipeline#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

