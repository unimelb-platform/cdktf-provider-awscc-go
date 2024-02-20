package codestarnotificationsnotificationrule


type CodestarnotificationsNotificationRuleTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarnotifications_notification_rule#target_address CodestarnotificationsNotificationRule#target_address}.
	TargetAddress *string `field:"required" json:"targetAddress" yaml:"targetAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarnotifications_notification_rule#target_type CodestarnotificationsNotificationRule#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

