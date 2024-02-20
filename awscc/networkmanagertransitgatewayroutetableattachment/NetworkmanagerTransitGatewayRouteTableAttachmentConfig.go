package networkmanagertransitgatewayroutetableattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerTransitGatewayRouteTableAttachmentConfig struct {
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
	// The Id of peering between transit gateway and core network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#peering_id NetworkmanagerTransitGatewayRouteTableAttachment#peering_id}
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
	// The Arn of transit gateway route table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#transit_gateway_route_table_arn NetworkmanagerTransitGatewayRouteTableAttachment#transit_gateway_route_table_arn}
	TransitGatewayRouteTableArn *string `field:"required" json:"transitGatewayRouteTableArn" yaml:"transitGatewayRouteTableArn"`
	// The attachment to move from one segment to another.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#proposed_segment_change NetworkmanagerTransitGatewayRouteTableAttachment#proposed_segment_change}
	ProposedSegmentChange *NetworkmanagerTransitGatewayRouteTableAttachmentProposedSegmentChange `field:"optional" json:"proposedSegmentChange" yaml:"proposedSegmentChange"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_transit_gateway_route_table_attachment#tags NetworkmanagerTransitGatewayRouteTableAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

