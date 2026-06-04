package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleConditionsHostHeaderConfig struct {
	// The host names.
	//
	// The maximum size of each name is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//  If you specify multiple strings, the condition is satisfied if one of the strings matches the host name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener_rule#values Elasticloadbalancingv2ListenerRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

