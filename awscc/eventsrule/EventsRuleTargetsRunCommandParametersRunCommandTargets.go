package eventsrule


type EventsRuleTargetsRunCommandParametersRunCommandTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/events_rule#key EventsRule#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/events_rule#values EventsRule#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

