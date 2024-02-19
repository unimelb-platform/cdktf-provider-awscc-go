package eventsrule


type EventsRuleTargetsSqsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#message_group_id EventsRule#message_group_id}.
	MessageGroupId *string `field:"required" json:"messageGroupId" yaml:"messageGroupId"`
}

