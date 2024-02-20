package ec2transitgatewaymulticastgroupsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayMulticastGroupSourceConfig struct {
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
	// The IP address assigned to the transit gateway multicast group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_multicast_group_source#group_ip_address Ec2TransitGatewayMulticastGroupSource#group_ip_address}
	GroupIpAddress *string `field:"required" json:"groupIpAddress" yaml:"groupIpAddress"`
	// The ID of the transit gateway attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_multicast_group_source#network_interface_id Ec2TransitGatewayMulticastGroupSource#network_interface_id}
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The ID of the transit gateway multicast domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_multicast_group_source#transit_gateway_multicast_domain_id Ec2TransitGatewayMulticastGroupSource#transit_gateway_multicast_domain_id}
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

