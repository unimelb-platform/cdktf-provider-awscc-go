package networkmanagerdirectconnectgatewayattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerDirectConnectGatewayAttachmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of a core network for the Direct Connect Gateway attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment#core_network_id NetworkmanagerDirectConnectGatewayAttachment#core_network_id}
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The ARN of the Direct Connect Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment#direct_connect_gateway_arn NetworkmanagerDirectConnectGatewayAttachment#direct_connect_gateway_arn}
	DirectConnectGatewayArn *string `field:"required" json:"directConnectGatewayArn" yaml:"directConnectGatewayArn"`
	// The Regions where the edges are located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment#edge_locations NetworkmanagerDirectConnectGatewayAttachment#edge_locations}
	EdgeLocations *[]*string `field:"required" json:"edgeLocations" yaml:"edgeLocations"`
	// The attachment to move from one network function group to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment#proposed_network_function_group_change NetworkmanagerDirectConnectGatewayAttachment#proposed_network_function_group_change}
	ProposedNetworkFunctionGroupChange *NetworkmanagerDirectConnectGatewayAttachmentProposedNetworkFunctionGroupChange `field:"optional" json:"proposedNetworkFunctionGroupChange" yaml:"proposedNetworkFunctionGroupChange"`
	// The attachment to move from one segment to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment#proposed_segment_change NetworkmanagerDirectConnectGatewayAttachment#proposed_segment_change}
	ProposedSegmentChange *NetworkmanagerDirectConnectGatewayAttachmentProposedSegmentChange `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// Tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkmanager_direct_connect_gateway_attachment#tags NetworkmanagerDirectConnectGatewayAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

