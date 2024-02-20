package ec2vpnconnectionroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpnConnectionRouteConfig struct {
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
	// The CIDR block associated with the local subnet of the customer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_connection_route#destination_cidr_block Ec2VpnConnectionRoute#destination_cidr_block}
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_connection_route#vpn_connection_id Ec2VpnConnectionRoute#vpn_connection_id}
	VpnConnectionId *string `field:"required" json:"vpnConnectionId" yaml:"vpnConnectionId"`
}

