package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConfig struct {
	// The baseline constraints resource for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#constraints_resource SagemakerMonitoringSchedule#constraints_resource}
	ConstraintsResource *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConfigConstraintsResource `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// The baseline statistics resource for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#statistics_resource SagemakerMonitoringSchedule#statistics_resource}
	StatisticsResource *SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionBaselineConfigStatisticsResource `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}

