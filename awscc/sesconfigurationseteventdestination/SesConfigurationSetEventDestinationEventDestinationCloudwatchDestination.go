package sesconfigurationseteventdestination


type SesConfigurationSetEventDestinationEventDestinationCloudwatchDestination struct {
	// A list of dimensions upon which to categorize your emails when you publish email sending events to Amazon CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set_event_destination#dimension_configurations SesConfigurationSetEventDestination#dimension_configurations}
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
}

