package quicksighttopic


type QuicksightTopicDataSetsFiltersNumericEqualityFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#aggregation QuicksightTopic#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#constant QuicksightTopic#constant}.
	Constant *QuicksightTopicDataSetsFiltersNumericEqualityFilterConstant `field:"optional" json:"constant" yaml:"constant"`
}

