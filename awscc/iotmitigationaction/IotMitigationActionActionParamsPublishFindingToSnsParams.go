package iotmitigationaction


type IotMitigationActionActionParamsPublishFindingToSnsParams struct {
	// The ARN of the topic to which you want to publish the findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#topic_arn IotMitigationAction#topic_arn}
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

