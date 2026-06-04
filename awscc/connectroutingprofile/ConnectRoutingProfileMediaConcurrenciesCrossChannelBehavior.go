package connectroutingprofile


type ConnectRoutingProfileMediaConcurrenciesCrossChannelBehavior struct {
	// Specifies the other channels that can be routed to an agent handling their current channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_routing_profile#behavior_type ConnectRoutingProfile#behavior_type}
	BehaviorType *string `field:"optional" json:"behaviorType" yaml:"behaviorType"`
}

