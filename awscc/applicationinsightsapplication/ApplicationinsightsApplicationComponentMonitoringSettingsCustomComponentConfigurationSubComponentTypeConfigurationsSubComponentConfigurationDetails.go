package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails struct {
	// A list of metrics to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#alarm_metrics ApplicationinsightsApplication#alarm_metrics}
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of logs to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#logs ApplicationinsightsApplication#logs}
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of Windows Events to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#windows_events ApplicationinsightsApplication#windows_events}
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

