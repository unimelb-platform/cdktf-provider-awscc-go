package appintegrationseventintegration


type AppintegrationsEventIntegrationEventFilter struct {
	// The source of the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appintegrations_event_integration#source AppintegrationsEventIntegration#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

