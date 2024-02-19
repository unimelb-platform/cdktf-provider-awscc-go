package appflowconnector


type AppflowConnectorConnectorProvisioningConfig struct {
	// Contains information about the configuration of the lambda which is being registered as the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector#lambda AppflowConnector#lambda}
	Lambda *AppflowConnectorConnectorProvisioningConfigLambda `field:"optional" json:"lambda" yaml:"lambda"`
}

