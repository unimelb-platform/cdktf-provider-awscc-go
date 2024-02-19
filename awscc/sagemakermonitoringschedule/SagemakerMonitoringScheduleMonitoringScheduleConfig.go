package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfig struct {
	// Defines the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#monitoring_job_definition SagemakerMonitoringSchedule#monitoring_job_definition}
	MonitoringJobDefinition *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinition `field:"optional" json:"monitoringJobDefinition" yaml:"monitoringJobDefinition"`
	// Name of the job definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#monitoring_job_definition_name SagemakerMonitoringSchedule#monitoring_job_definition_name}
	MonitoringJobDefinitionName *string `field:"optional" json:"monitoringJobDefinitionName" yaml:"monitoringJobDefinitionName"`
	// The type of monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#monitoring_type SagemakerMonitoringSchedule#monitoring_type}
	MonitoringType *string `field:"optional" json:"monitoringType" yaml:"monitoringType"`
	// Configuration details about the monitoring schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#schedule_config SagemakerMonitoringSchedule#schedule_config}
	ScheduleConfig *SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
}

