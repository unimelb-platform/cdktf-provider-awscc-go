package eventsrule


type EventsRuleTargetsBatchParametersRetryStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#attempts EventsRule#attempts}.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
}

