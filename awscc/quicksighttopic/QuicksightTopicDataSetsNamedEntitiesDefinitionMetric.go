package quicksighttopic


type QuicksightTopicDataSetsNamedEntitiesDefinitionMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#aggregation QuicksightTopic#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#aggregation_function_parameters QuicksightTopic#aggregation_function_parameters}.
	AggregationFunctionParameters *map[string]*string `field:"optional" json:"aggregationFunctionParameters" yaml:"aggregationFunctionParameters"`
}

