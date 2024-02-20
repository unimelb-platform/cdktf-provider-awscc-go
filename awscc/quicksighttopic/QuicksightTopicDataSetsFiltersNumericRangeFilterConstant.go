package quicksighttopic


type QuicksightTopicDataSetsFiltersNumericRangeFilterConstant struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#constant_type QuicksightTopic#constant_type}.
	ConstantType *string `field:"optional" json:"constantType" yaml:"constantType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#range_constant QuicksightTopic#range_constant}.
	RangeConstant *QuicksightTopicDataSetsFiltersNumericRangeFilterConstantRangeConstant `field:"optional" json:"rangeConstant" yaml:"rangeConstant"`
}

