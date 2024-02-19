package quicksighttopic


type QuicksightTopicDataSetsFiltersNumericRangeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#aggregation QuicksightTopic#aggregation}.
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#constant QuicksightTopic#constant}.
	Constant *QuicksightTopicDataSetsFiltersNumericRangeFilterConstant `field:"optional" json:"constant" yaml:"constant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#inclusive QuicksightTopic#inclusive}.
	Inclusive interface{} `field:"optional" json:"inclusive" yaml:"inclusive"`
}

