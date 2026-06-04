package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurations struct {
	// The configuration settings of sub components.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationinsights_application#sub_component_configuration_details ApplicationinsightsApplication#sub_component_configuration_details}
	SubComponentConfigurationDetails *ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationSubComponentTypeConfigurationsSubComponentConfigurationDetails `field:"optional" json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub component type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationinsights_application#sub_component_type ApplicationinsightsApplication#sub_component_type}
	SubComponentType *string `field:"optional" json:"subComponentType" yaml:"subComponentType"`
}

