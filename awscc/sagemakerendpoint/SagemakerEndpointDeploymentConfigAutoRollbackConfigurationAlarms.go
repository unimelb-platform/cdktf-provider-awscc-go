package sagemakerendpoint


type SagemakerEndpointDeploymentConfigAutoRollbackConfigurationAlarms struct {
	// The name of the CloudWatch alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_endpoint#alarm_name SagemakerEndpoint#alarm_name}
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
}

