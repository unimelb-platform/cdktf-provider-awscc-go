package eventsrule


type EventsRuleTargetsKinesisParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#partition_key_path EventsRule#partition_key_path}.
	PartitionKeyPath *string `field:"required" json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

