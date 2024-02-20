package eventsendpoint


type EventsEndpointRoutingConfigFailoverConfigPrimary struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_endpoint#health_check EventsEndpoint#health_check}.
	HealthCheck *string `field:"required" json:"healthCheck" yaml:"healthCheck"`
}

