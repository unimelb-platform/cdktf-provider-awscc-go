package lambdaeventsourcemapping


type LambdaEventSourceMappingFilterCriteriaFilters struct {
	// The filter pattern that defines which events should be passed for invocations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#pattern LambdaEventSourceMapping#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

