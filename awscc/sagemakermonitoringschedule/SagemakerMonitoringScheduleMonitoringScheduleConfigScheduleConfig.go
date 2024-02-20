package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigScheduleConfig struct {
	// A cron expression or 'NOW' that describes details about the monitoring schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#schedule_expression SagemakerMonitoringSchedule#schedule_expression}
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Data Analysis end time, e.g. PT0H.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#data_analysis_end_time SagemakerMonitoringSchedule#data_analysis_end_time}
	DataAnalysisEndTime *string `field:"optional" json:"dataAnalysisEndTime" yaml:"dataAnalysisEndTime"`
	// Data Analysis start time, e.g. -PT1H.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#data_analysis_start_time SagemakerMonitoringSchedule#data_analysis_start_time}
	DataAnalysisStartTime *string `field:"optional" json:"dataAnalysisStartTime" yaml:"dataAnalysisStartTime"`
}

