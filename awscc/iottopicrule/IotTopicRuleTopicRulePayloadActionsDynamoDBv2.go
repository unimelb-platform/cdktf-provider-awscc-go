package iottopicrule


type IotTopicRuleTopicRulePayloadActionsDynamoDBv2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#put_item IotTopicRule#put_item}.
	PutItem *IotTopicRuleTopicRulePayloadActionsDynamoDBv2PutItem `field:"optional" json:"putItem" yaml:"putItem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

