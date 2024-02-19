package connectroutingprofile


type ConnectRoutingProfileQueueConfigs struct {
	// The delay, in seconds, a contact should wait in the queue before they are routed to an available agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_routing_profile#delay ConnectRoutingProfile#delay}
	Delay *float64 `field:"required" json:"delay" yaml:"delay"`
	// The order in which contacts are to be handled for the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_routing_profile#priority ConnectRoutingProfile#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Contains the channel and queue identifier for a routing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_routing_profile#queue_reference ConnectRoutingProfile#queue_reference}
	QueueReference *ConnectRoutingProfileQueueConfigsQueueReference `field:"required" json:"queueReference" yaml:"queueReference"`
}

