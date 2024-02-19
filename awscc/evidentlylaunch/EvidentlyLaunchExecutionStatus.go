package evidentlylaunch


type EvidentlyLaunchExecutionStatus struct {
	// Provide START or STOP action to apply on a launch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#status EvidentlyLaunch#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Provide CANCELLED or COMPLETED as the launch desired state. Defaults to Completed if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#desired_state EvidentlyLaunch#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Provide a reason for stopping the launch. Defaults to empty if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#reason EvidentlyLaunch#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

