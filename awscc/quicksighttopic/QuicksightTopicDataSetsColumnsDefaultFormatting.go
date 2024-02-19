package quicksighttopic


type QuicksightTopicDataSetsColumnsDefaultFormatting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#display_format QuicksightTopic#display_format}.
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_topic#display_format_options QuicksightTopic#display_format_options}.
	DisplayFormatOptions *QuicksightTopicDataSetsColumnsDefaultFormattingDisplayFormatOptions `field:"optional" json:"displayFormatOptions" yaml:"displayFormatOptions"`
}

