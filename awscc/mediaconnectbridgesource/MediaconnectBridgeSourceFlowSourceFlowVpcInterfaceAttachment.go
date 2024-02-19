package mediaconnectbridgesource


type MediaconnectBridgeSourceFlowSourceFlowVpcInterfaceAttachment struct {
	// The name of the VPC interface to use for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#vpc_interface_name MediaconnectBridgeSource#vpc_interface_name}
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

