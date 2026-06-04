package medialiveeventbridgeruletemplate


type MedialiveEventBridgeRuleTemplateEventTargets struct {
	// Target ARNs must be either an SNS topic or CloudWatch log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/medialive_event_bridge_rule_template#arn MedialiveEventBridgeRuleTemplate#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

