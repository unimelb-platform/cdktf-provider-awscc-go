package apprunnerservice


type ApprunnerServiceObservabilityConfiguration struct {
	// Observability enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#observability_enabled ApprunnerService#observability_enabled}
	ObservabilityEnabled interface{} `field:"required" json:"observabilityEnabled" yaml:"observabilityEnabled"`
	// The Amazon Resource Name (ARN) of the App Runner ObservabilityConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#observability_configuration_arn ApprunnerService#observability_configuration_arn}
	ObservabilityConfigurationArn *string `field:"optional" json:"observabilityConfigurationArn" yaml:"observabilityConfigurationArn"`
}

