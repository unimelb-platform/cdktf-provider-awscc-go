package datazoneenvironmentblueprintconfiguration


type DatazoneEnvironmentBlueprintConfigurationRegionalParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#parameters DatazoneEnvironmentBlueprintConfiguration#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#region DatazoneEnvironmentBlueprintConfiguration#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

