package iottopicrule


type IotTopicRuleTopicRulePayloadActionsHttpAuthSigv4 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#service_name IotTopicRule#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#signing_region IotTopicRule#signing_region}.
	SigningRegion *string `field:"required" json:"signingRegion" yaml:"signingRegion"`
}

