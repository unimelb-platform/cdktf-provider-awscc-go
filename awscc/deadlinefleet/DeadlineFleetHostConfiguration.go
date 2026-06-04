package deadlinefleet


type DeadlineFleetHostConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#script_body DeadlineFleet#script_body}.
	ScriptBody *string `field:"optional" json:"scriptBody" yaml:"scriptBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#script_timeout_seconds DeadlineFleet#script_timeout_seconds}.
	ScriptTimeoutSeconds *float64 `field:"optional" json:"scriptTimeoutSeconds" yaml:"scriptTimeoutSeconds"`
}

