package eventsendpoint


type EventsEndpointRoutingConfigFailoverConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_endpoint#primary EventsEndpoint#primary}.
	Primary *EventsEndpointRoutingConfigFailoverConfigPrimary `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_endpoint#secondary EventsEndpoint#secondary}.
	Secondary *EventsEndpointRoutingConfigFailoverConfigSecondary `field:"required" json:"secondary" yaml:"secondary"`
}

