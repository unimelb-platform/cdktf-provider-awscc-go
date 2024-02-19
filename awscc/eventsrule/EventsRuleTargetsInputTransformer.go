package eventsrule


type EventsRuleTargetsInputTransformer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#input_template EventsRule#input_template}.
	InputTemplate *string `field:"required" json:"inputTemplate" yaml:"inputTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_rule#input_paths_map EventsRule#input_paths_map}.
	InputPathsMap *map[string]*string `field:"optional" json:"inputPathsMap" yaml:"inputPathsMap"`
}

