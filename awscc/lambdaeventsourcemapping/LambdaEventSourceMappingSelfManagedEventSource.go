package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSource struct {
	// The endpoints for a self-managed event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}
	Endpoints *LambdaEventSourceMappingSelfManagedEventSourceEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
}

