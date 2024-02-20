package eventsrule


type EventsRuleTargetsRunCommandParametersRunCommandTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#key EventsRule#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#values EventsRule#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

