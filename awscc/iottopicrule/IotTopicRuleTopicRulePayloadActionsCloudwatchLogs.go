package iottopicrule


type IotTopicRuleTopicRulePayloadActionsCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#log_group_name IotTopicRule#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#batch_mode IotTopicRule#batch_mode}.
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
}

