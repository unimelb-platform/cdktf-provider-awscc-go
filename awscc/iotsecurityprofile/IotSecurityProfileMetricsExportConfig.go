package iotsecurityprofile


type IotSecurityProfileMetricsExportConfig struct {
	// The topic for metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#mqtt_topic IotSecurityProfile#mqtt_topic}
	MqttTopic *string `field:"required" json:"mqttTopic" yaml:"mqttTopic"`
	// The ARN of the role that grants permission to publish to mqtt topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#role_arn IotSecurityProfile#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

