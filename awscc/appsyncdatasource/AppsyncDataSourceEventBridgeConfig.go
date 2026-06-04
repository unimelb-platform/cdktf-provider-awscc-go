package appsyncdatasource


type AppsyncDataSourceEventBridgeConfig struct {
	// ARN for the EventBridge bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_data_source#event_bus_arn AppsyncDataSource#event_bus_arn}
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
}

