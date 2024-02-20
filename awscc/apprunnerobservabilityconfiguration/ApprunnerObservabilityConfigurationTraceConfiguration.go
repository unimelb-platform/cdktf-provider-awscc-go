package apprunnerobservabilityconfiguration


type ApprunnerObservabilityConfigurationTraceConfiguration struct {
	// The implementation provider chosen for tracing App Runner services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_observability_configuration#vendor ApprunnerObservabilityConfiguration#vendor}
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
}

