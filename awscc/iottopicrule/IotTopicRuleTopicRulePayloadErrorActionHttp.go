package iottopicrule


type IotTopicRuleTopicRulePayloadErrorActionHttp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#auth IotTopicRule#auth}.
	Auth *IotTopicRuleTopicRulePayloadErrorActionHttpAuth `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#confirmation_url IotTopicRule#confirmation_url}.
	ConfirmationUrl *string `field:"optional" json:"confirmationUrl" yaml:"confirmationUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#headers IotTopicRule#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_topic_rule#url IotTopicRule#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

