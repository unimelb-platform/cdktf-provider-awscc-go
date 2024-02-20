package iottopicrule


type IotTopicRuleTopicRulePayloadErrorActionLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#device_id IotTopicRule#device_id}.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#latitude IotTopicRule#latitude}.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#longitude IotTopicRule#longitude}.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#tracker_name IotTopicRule#tracker_name}.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#timestamp IotTopicRule#timestamp}.
	Timestamp *IotTopicRuleTopicRulePayloadErrorActionLocationTimestamp `field:"optional" json:"timestamp" yaml:"timestamp"`
}

