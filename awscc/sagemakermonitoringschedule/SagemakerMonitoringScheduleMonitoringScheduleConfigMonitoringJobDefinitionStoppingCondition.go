package sagemakermonitoringschedule


type SagemakerMonitoringScheduleMonitoringScheduleConfigMonitoringJobDefinitionStoppingCondition struct {
	// The maximum runtime allowed in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_monitoring_schedule#max_runtime_in_seconds SagemakerMonitoringSchedule#max_runtime_in_seconds}
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

