package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsAlarms struct {
	// The name of the CloudWatch alarm to be monitored for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#alarm_name ApplicationinsightsApplication#alarm_name}
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Indicates the degree of outage when the alarm goes off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#severity ApplicationinsightsApplication#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}

