package iotsecurityprofile


type IotSecurityProfileAlertTargets struct {
	// The ARN of the notification target to which alerts are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#alert_target_arn IotSecurityProfile#alert_target_arn}
	AlertTargetArn *string `field:"optional" json:"alertTargetArn" yaml:"alertTargetArn"`
	// The ARN of the role that grants permission to send alerts to the notification target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_security_profile#role_arn IotSecurityProfile#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

