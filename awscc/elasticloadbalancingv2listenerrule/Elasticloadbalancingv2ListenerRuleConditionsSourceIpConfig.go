package elasticloadbalancingv2listenerrule


type Elasticloadbalancingv2ListenerRuleConditionsSourceIpConfig struct {
	// The source IP addresses, in CIDR format.
	//
	// You can use both IPv4 and IPv6 addresses. Wildcards are not supported.
	//  If you specify multiple addresses, the condition is satisfied if the source IP address of the request matches one of the CIDR blocks. This condition is not satisfied by the addresses in the X-Forwarded-For header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener_rule#values Elasticloadbalancingv2ListenerRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

