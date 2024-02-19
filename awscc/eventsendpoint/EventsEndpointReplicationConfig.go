package eventsendpoint


type EventsEndpointReplicationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_endpoint#state EventsEndpoint#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

