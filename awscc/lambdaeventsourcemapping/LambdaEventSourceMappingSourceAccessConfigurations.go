package lambdaeventsourcemapping


type LambdaEventSourceMappingSourceAccessConfigurations struct {
	// The type of source access configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#type LambdaEventSourceMapping#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The URI for the source access configuration resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#uri LambdaEventSourceMapping#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

