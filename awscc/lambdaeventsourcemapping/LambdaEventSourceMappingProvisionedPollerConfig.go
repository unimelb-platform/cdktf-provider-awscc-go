package lambdaeventsourcemapping


type LambdaEventSourceMappingProvisionedPollerConfig struct {
	// The maximum number of event pollers this event source can scale up to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#maximum_pollers LambdaEventSourceMapping#maximum_pollers}
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// The minimum number of event pollers this event source can scale down to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#minimum_pollers LambdaEventSourceMapping#minimum_pollers}
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
}

