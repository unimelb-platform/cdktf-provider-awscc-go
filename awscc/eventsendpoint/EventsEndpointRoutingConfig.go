package eventsendpoint


type EventsEndpointRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_endpoint#failover_config EventsEndpoint#failover_config}.
	FailoverConfig *EventsEndpointRoutingConfigFailoverConfig `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}

