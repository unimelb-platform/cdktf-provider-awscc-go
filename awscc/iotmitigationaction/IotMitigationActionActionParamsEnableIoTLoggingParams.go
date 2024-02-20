package iotmitigationaction


type IotMitigationActionActionParamsEnableIoTLoggingParams struct {
	// Specifies which types of information are logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#log_level IotMitigationAction#log_level}
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
	// The ARN of the IAM role used for logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#role_arn_for_logging IotMitigationAction#role_arn_for_logging}
	RoleArnForLogging *string `field:"required" json:"roleArnForLogging" yaml:"roleArnForLogging"`
}

