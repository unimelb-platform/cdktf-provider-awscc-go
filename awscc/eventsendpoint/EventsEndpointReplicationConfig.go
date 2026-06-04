package eventsendpoint


type EventsEndpointReplicationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/events_endpoint#state EventsEndpoint#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

