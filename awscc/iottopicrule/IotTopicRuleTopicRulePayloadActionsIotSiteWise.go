package iottopicrule


type IotTopicRuleTopicRulePayloadActionsIotSiteWise struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#put_asset_property_value_entries IotTopicRule#put_asset_property_value_entries}.
	PutAssetPropertyValueEntries interface{} `field:"required" json:"putAssetPropertyValueEntries" yaml:"putAssetPropertyValueEntries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

