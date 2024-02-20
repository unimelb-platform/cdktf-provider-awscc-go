package iottopicrule


type IotTopicRuleTopicRulePayloadActionsIotSiteWisePutAssetPropertyValueEntriesPropertyValuesTimestamp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#time_in_seconds IotTopicRule#time_in_seconds}.
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#offset_in_nanos IotTopicRule#offset_in_nanos}.
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}

