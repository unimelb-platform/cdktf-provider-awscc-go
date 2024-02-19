package stepfunctionsstatemachine


type StepfunctionsStateMachineLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine#destinations StepfunctionsStateMachine#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine#include_execution_data StepfunctionsStateMachine#include_execution_data}.
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine#level StepfunctionsStateMachine#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
}

