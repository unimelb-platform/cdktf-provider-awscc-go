package eventsrule


type EventsRuleTargetsRunCommandParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#run_command_targets EventsRule#run_command_targets}.
	RunCommandTargets interface{} `field:"required" json:"runCommandTargets" yaml:"runCommandTargets"`
}

