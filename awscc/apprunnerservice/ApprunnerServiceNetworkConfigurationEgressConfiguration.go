package apprunnerservice


type ApprunnerServiceNetworkConfigurationEgressConfiguration struct {
	// Network egress type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#egress_type ApprunnerService#egress_type}
	EgressType *string `field:"required" json:"egressType" yaml:"egressType"`
	// The Amazon Resource Name (ARN) of the App Runner VpcConnector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#vpc_connector_arn ApprunnerService#vpc_connector_arn}
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

