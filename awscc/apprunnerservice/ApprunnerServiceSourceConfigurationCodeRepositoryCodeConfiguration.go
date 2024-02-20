package apprunnerservice


type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration struct {
	// Configuration Source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#configuration_source ApprunnerService#configuration_source}
	ConfigurationSource *string `field:"required" json:"configurationSource" yaml:"configurationSource"`
	// Code Configuration Values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#code_configuration_values ApprunnerService#code_configuration_values}
	CodeConfigurationValues *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues `field:"optional" json:"codeConfigurationValues" yaml:"codeConfigurationValues"`
}

