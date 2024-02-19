package iottopicrule


type IotTopicRuleTopicRulePayloadActionsHttp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#url IotTopicRule#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#auth IotTopicRule#auth}.
	Auth *IotTopicRuleTopicRulePayloadActionsHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#confirmation_url IotTopicRule#confirmation_url}.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_topic_rule#headers IotTopicRule#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
}

