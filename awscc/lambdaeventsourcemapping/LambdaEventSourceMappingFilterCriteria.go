package lambdaeventsourcemapping


type LambdaEventSourceMappingFilterCriteria struct {
	// List of filters of this FilterCriteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#filters LambdaEventSourceMapping#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

