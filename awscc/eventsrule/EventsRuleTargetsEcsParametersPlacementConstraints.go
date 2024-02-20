package eventsrule


type EventsRuleTargetsEcsParametersPlacementConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#expression EventsRule#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#type EventsRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

