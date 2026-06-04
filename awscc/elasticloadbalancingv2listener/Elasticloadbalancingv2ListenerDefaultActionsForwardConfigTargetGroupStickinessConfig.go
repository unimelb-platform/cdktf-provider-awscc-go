package elasticloadbalancingv2listener


type Elasticloadbalancingv2ListenerDefaultActionsForwardConfigTargetGroupStickinessConfig struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group.
	//
	// The range is 1-604800 seconds (7 days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener#duration_seconds Elasticloadbalancingv2Listener#duration_seconds}
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Indicates whether target group stickiness is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener#enabled Elasticloadbalancingv2Listener#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

