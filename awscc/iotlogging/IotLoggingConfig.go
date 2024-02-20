package iotlogging

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotLoggingConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_logging#account_id IotLogging#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_logging#default_log_level IotLogging#default_log_level}
	DefaultLogLevel *string `field:"required" json:"defaultLogLevel" yaml:"defaultLogLevel"`
	// The ARN of the role that allows IoT to write to Cloudwatch logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_logging#role_arn IotLogging#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

