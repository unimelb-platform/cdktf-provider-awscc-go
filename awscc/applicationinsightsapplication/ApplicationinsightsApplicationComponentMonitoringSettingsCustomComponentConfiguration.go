package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfiguration struct {
	// The configuration settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#configuration_details ApplicationinsightsApplication#configuration_details}
	ConfigurationDetails *ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetails `field:"optional" json:"configurationDetails" yaml:"configurationDetails"`
	// Sub component configurations of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#sub_component_type_configurations ApplicationinsightsApplication#sub_component_type_configurations}
	SubComponentTypeConfigurations interface{} `field:"optional" json:"subComponentTypeConfigurations" yaml:"subComponentTypeConfigurations"`
}

