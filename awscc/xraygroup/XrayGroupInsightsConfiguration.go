package xraygroup


type XrayGroupInsightsConfiguration struct {
	// Set the InsightsEnabled value to true to enable insights or false to disable insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#insights_enabled XrayGroup#insights_enabled}
	InsightsEnabled interface{} `field:"optional" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Set the NotificationsEnabled value to true to enable insights notifications.
	//
	// Notifications can only be enabled on a group with InsightsEnabled set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#notifications_enabled XrayGroup#notifications_enabled}
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

