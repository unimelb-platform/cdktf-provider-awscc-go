package lambdaeventsourcemapping


type LambdaEventSourceMappingDestinationConfigOnFailure struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#destination LambdaEventSourceMapping#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
}

