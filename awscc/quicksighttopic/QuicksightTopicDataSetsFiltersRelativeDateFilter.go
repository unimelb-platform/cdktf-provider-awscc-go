package quicksighttopic


type QuicksightTopicDataSetsFiltersRelativeDateFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#constant QuicksightTopic#constant}.
	Constant *QuicksightTopicDataSetsFiltersRelativeDateFilterConstant `field:"optional" json:"constant" yaml:"constant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#relative_date_filter_function QuicksightTopic#relative_date_filter_function}.
	RelativeDateFilterFunction *string `field:"optional" json:"relativeDateFilterFunction" yaml:"relativeDateFilterFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#time_granularity QuicksightTopic#time_granularity}.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

