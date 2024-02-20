package iottopicrule


type IotTopicRuleTopicRulePayloadActionsRepublish struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#topic IotTopicRule#topic}.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#headers IotTopicRule#headers}.
	Headers *IotTopicRuleTopicRulePayloadActionsRepublishHeaders `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#qos IotTopicRule#qos}.
	Qos *float64 `field:"optional" json:"qos" yaml:"qos"`
}

