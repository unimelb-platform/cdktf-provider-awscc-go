package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetailsProcesses struct {
	// A list of metrics to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationinsights_application#alarm_metrics ApplicationinsightsApplication#alarm_metrics}
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// The name of the process to be monitored for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationinsights_application#process_name ApplicationinsightsApplication#process_name}
	ProcessName *string `field:"optional" json:"processName" yaml:"processName"`
}

