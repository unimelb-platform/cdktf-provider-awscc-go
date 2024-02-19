package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetails struct {
	// A list of metrics to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#alarm_metrics ApplicationinsightsApplication#alarm_metrics}
	AlarmMetrics interface{} `field:"optional" json:"alarmMetrics" yaml:"alarmMetrics"`
	// A list of alarms to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#alarms ApplicationinsightsApplication#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// The HA cluster Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#ha_cluster_prometheus_exporter ApplicationinsightsApplication#ha_cluster_prometheus_exporter}
	HaClusterPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsHaClusterPrometheusExporter `field:"optional" json:"haClusterPrometheusExporter" yaml:"haClusterPrometheusExporter"`
	// The HANA DB Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#hana_prometheus_exporter ApplicationinsightsApplication#hana_prometheus_exporter}
	HanaPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsHanaPrometheusExporter `field:"optional" json:"hanaPrometheusExporter" yaml:"hanaPrometheusExporter"`
	// The JMX Prometheus Exporter settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#jmx_prometheus_exporter ApplicationinsightsApplication#jmx_prometheus_exporter}
	JmxPrometheusExporter *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsJmxPrometheusExporter `field:"optional" json:"jmxPrometheusExporter" yaml:"jmxPrometheusExporter"`
	// A list of logs to monitor for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#logs ApplicationinsightsApplication#logs}
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// A list of Windows Events to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#windows_events ApplicationinsightsApplication#windows_events}
	WindowsEvents interface{} `field:"optional" json:"windowsEvents" yaml:"windowsEvents"`
}

