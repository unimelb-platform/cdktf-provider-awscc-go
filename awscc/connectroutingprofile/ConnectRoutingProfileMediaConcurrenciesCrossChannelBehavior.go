package connectroutingprofile


type ConnectRoutingProfileMediaConcurrenciesCrossChannelBehavior struct {
	// Specifies the other channels that can be routed to an agent handling their current channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_routing_profile#behavior_type ConnectRoutingProfile#behavior_type}
	BehaviorType *string `field:"required" json:"behaviorType" yaml:"behaviorType"`
}

