package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfiguration struct {
	// The verbosity of the CloudWatch Logs for an application. You can set it to `INFO`, `WARN`, `ERROR`, or `DEBUG`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#log_level Kinesisanalyticsv2Application#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

