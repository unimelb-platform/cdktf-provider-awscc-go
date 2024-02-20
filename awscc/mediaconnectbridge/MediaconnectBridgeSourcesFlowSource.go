package mediaconnectbridge


type MediaconnectBridgeSourcesFlowSource struct {
	// The ARN of the cloud flow used as a source of this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#flow_arn MediaconnectBridge#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the flow source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the VPC interface attachment to use for this source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#flow_vpc_interface_attachment MediaconnectBridge#flow_vpc_interface_attachment}
	FlowVpcInterfaceAttachment *MediaconnectBridgeSourcesFlowSourceFlowVpcInterfaceAttachment `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
}

