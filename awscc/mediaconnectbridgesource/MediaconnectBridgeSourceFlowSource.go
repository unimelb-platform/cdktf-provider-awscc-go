package mediaconnectbridgesource


type MediaconnectBridgeSourceFlowSource struct {
	// The ARN of the cloud flow used as a source of this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#flow_arn MediaconnectBridgeSource#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the VPC interface attachment to use for this source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#flow_vpc_interface_attachment MediaconnectBridgeSource#flow_vpc_interface_attachment}
	FlowVpcInterfaceAttachment *MediaconnectBridgeSourceFlowSourceFlowVpcInterfaceAttachment `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
}

