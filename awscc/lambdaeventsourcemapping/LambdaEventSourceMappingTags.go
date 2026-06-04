package lambdaeventsourcemapping


type LambdaEventSourceMappingTags struct {
	// The key for this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#key LambdaEventSourceMapping#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#value LambdaEventSourceMapping#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

