package appconfigenvironment


type AppconfigEnvironmentMonitors struct {
	// Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appconfig_environment#alarm_arn AppconfigEnvironment#alarm_arn}
	AlarmArn *string `field:"optional" json:"alarmArn" yaml:"alarmArn"`
	// ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor AlarmArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appconfig_environment#alarm_role_arn AppconfigEnvironment#alarm_role_arn}
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

