package stepfunctionsstatemachine


type StepfunctionsStateMachineLoggingConfigurationDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine#cloudwatch_logs_log_group StepfunctionsStateMachine#cloudwatch_logs_log_group}.
	CloudwatchLogsLogGroup *StepfunctionsStateMachineLoggingConfigurationDestinationsCloudwatchLogsLogGroup `field:"optional" json:"cloudwatchLogsLogGroup" yaml:"cloudwatchLogsLogGroup"`
}

