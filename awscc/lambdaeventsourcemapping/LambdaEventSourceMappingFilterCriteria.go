package lambdaeventsourcemapping


type LambdaEventSourceMappingFilterCriteria struct {
	// A list of filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#filters LambdaEventSourceMapping#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

