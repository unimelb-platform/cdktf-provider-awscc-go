package eventsendpoint


type EventsEndpointEventBuses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_endpoint#event_bus_arn EventsEndpoint#event_bus_arn}.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

