package lambdaeventsourcemapping


type LambdaEventSourceMappingDestinationConfig struct {
	// The destination configuration for failed invocations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#on_failure LambdaEventSourceMapping#on_failure}
	OnFailure *LambdaEventSourceMappingDestinationConfigOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
}

