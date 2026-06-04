package iottopicrule


type IotTopicRuleTopicRulePayloadErrorActionSqs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#queue_url IotTopicRule#queue_url}.
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#use_base_64 IotTopicRule#use_base_64}.
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

