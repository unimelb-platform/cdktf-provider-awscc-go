package iotaccountauditconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotAccountAuditConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_account_audit_configuration#account_id IotAccountAuditConfiguration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Specifies which audit checks are enabled and disabled for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_account_audit_configuration#audit_check_configurations IotAccountAuditConfiguration#audit_check_configurations}
	AuditCheckConfigurations *IotAccountAuditConfigurationAuditCheckConfigurations `field:"required" json:"auditCheckConfigurations" yaml:"auditCheckConfigurations"`
	// The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_account_audit_configuration#role_arn IotAccountAuditConfiguration#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Information about the targets to which audit notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_account_audit_configuration#audit_notification_target_configurations IotAccountAuditConfiguration#audit_notification_target_configurations}
	AuditNotificationTargetConfigurations *IotAccountAuditConfigurationAuditNotificationTargetConfigurations `field:"optional" json:"auditNotificationTargetConfigurations" yaml:"auditNotificationTargetConfigurations"`
}

