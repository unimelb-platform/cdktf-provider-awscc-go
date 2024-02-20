package iottopicrule


type IotTopicRuleTopicRulePayloadActionsFirehose struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#delivery_stream_name IotTopicRule#delivery_stream_name}.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#batch_mode IotTopicRule#batch_mode}.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#separator IotTopicRule#separator}.
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

