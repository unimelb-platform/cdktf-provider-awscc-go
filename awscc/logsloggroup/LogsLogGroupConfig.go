package logsloggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsLogGroupConfig struct {
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
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_group#data_protection_policy LogsLogGroup#data_protection_policy}
	DataProtectionPolicy *string `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_group#kms_key_id LogsLogGroup#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The class of the log group. Possible values are: STANDARD and INFREQUENT_ACCESS, with STANDARD being the default class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_group#log_group_class LogsLogGroup#log_group_class}
	LogGroupClass *string `field:"optional" json:"logGroupClass" yaml:"logGroupClass"`
	// The name of the log group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_group#log_group_name LogsLogGroup#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The number of days to retain the log events in the specified log group.
	//
	// Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, and 3653.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_group#retention_in_days LogsLogGroup#retention_in_days}
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_log_group#tags LogsLogGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

