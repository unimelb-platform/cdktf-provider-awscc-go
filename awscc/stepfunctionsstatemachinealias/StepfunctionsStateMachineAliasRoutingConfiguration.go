package stepfunctionsstatemachinealias


type StepfunctionsStateMachineAliasRoutingConfiguration struct {
	// The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_alias#state_machine_version_arn StepfunctionsStateMachineAlias#state_machine_version_arn}
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// The percentage of traffic you want to route to the state machine version.
	//
	// The sum of the weights in the routing configuration must be equal to 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_alias#weight StepfunctionsStateMachineAlias#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

