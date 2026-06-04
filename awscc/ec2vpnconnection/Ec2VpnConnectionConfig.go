package ec2vpnconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpnConnectionConfig struct {
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
	// The ID of the customer gateway at your end of the VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#customer_gateway_id Ec2VpnConnection#customer_gateway_id}
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// The type of VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#type Ec2VpnConnection#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Indicate whether to enable acceleration for the VPN connection.  Default: ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#enable_acceleration Ec2VpnConnection#enable_acceleration}
	EnableAcceleration interface{} `field:"optional" json:"enableAcceleration" yaml:"enableAcceleration"`
	// The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.  Default: ``0.0.0.0/0``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#local_ipv_4_network_cidr Ec2VpnConnection#local_ipv_4_network_cidr}
	LocalIpv4NetworkCidr *string `field:"optional" json:"localIpv4NetworkCidr" yaml:"localIpv4NetworkCidr"`
	// The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.  Default: ``::/0``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#local_ipv_6_network_cidr Ec2VpnConnection#local_ipv_6_network_cidr}
	LocalIpv6NetworkCidr *string `field:"optional" json:"localIpv6NetworkCidr" yaml:"localIpv6NetworkCidr"`
	// The type of IPv4 address assigned to the outside interface of the customer gateway device.
	//
	// Valid values: ``PrivateIpv4`` | ``PublicIpv4``
	//  Default: ``PublicIpv4``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#outside_ip_address_type Ec2VpnConnection#outside_ip_address_type}
	OutsideIpAddressType *string `field:"optional" json:"outsideIpAddressType" yaml:"outsideIpAddressType"`
	// The IPv4 CIDR on the AWS side of the VPN connection.  Default: ``0.0.0.0/0``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#remote_ipv_4_network_cidr Ec2VpnConnection#remote_ipv_4_network_cidr}
	RemoteIpv4NetworkCidr *string `field:"optional" json:"remoteIpv4NetworkCidr" yaml:"remoteIpv4NetworkCidr"`
	// The IPv6 CIDR on the AWS side of the VPN connection.  Default: ``::/0``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#remote_ipv_6_network_cidr Ec2VpnConnection#remote_ipv_6_network_cidr}
	RemoteIpv6NetworkCidr *string `field:"optional" json:"remoteIpv6NetworkCidr" yaml:"remoteIpv6NetworkCidr"`
	// Indicates whether the VPN connection uses static routes only.
	//
	// Static routes must be used for devices that don't support BGP.
	//  If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ``true``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#static_routes_only Ec2VpnConnection#static_routes_only}
	StaticRoutesOnly interface{} `field:"optional" json:"staticRoutesOnly" yaml:"staticRoutesOnly"`
	// Any tags assigned to the VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#tags Ec2VpnConnection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the transit gateway associated with the VPN connection.
	//
	// You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#transit_gateway_id Ec2VpnConnection#transit_gateway_id}
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The transit gateway attachment ID to use for the VPN tunnel.  Required if ``OutsideIpAddressType`` is set to ``PrivateIpv4``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#transport_transit_gateway_attachment_id Ec2VpnConnection#transport_transit_gateway_attachment_id}
	TransportTransitGatewayAttachmentId *string `field:"optional" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// Indicate whether the VPN tunnels process IPv4 or IPv6 traffic.  Default: ``ipv4``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#tunnel_inside_ip_version Ec2VpnConnection#tunnel_inside_ip_version}
	TunnelInsideIpVersion *string `field:"optional" json:"tunnelInsideIpVersion" yaml:"tunnelInsideIpVersion"`
	// The ID of the virtual private gateway at the AWS side of the VPN connection.
	//
	// You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#vpn_gateway_id Ec2VpnConnection#vpn_gateway_id}
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
	// The tunnel options for the VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#vpn_tunnel_options_specifications Ec2VpnConnection#vpn_tunnel_options_specifications}
	VpnTunnelOptionsSpecifications interface{} `field:"optional" json:"vpnTunnelOptionsSpecifications" yaml:"vpnTunnelOptionsSpecifications"`
}

