package sagemakerendpoint


type SagemakerEndpointTags struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_endpoint#key SagemakerEndpoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_endpoint#value SagemakerEndpoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

