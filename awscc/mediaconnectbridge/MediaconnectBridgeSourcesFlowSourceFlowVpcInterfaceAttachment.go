package mediaconnectbridge


type MediaconnectBridgeSourcesFlowSourceFlowVpcInterfaceAttachment struct {
	// The name of the VPC interface to use for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#vpc_interface_name MediaconnectBridge#vpc_interface_name}
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

