package connectroutingprofile


type ConnectRoutingProfileQueueConfigsQueueReference struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_routing_profile#channel ConnectRoutingProfile#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The Amazon Resource Name (ARN) for the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_routing_profile#queue_arn ConnectRoutingProfile#queue_arn}
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
}

