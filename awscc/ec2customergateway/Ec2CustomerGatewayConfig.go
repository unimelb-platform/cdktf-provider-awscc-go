package ec2customergateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2CustomerGatewayConfig struct {
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
	// The internet-routable IP address for the customer gateway's outside interface. The address must be static.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#ip_address Ec2CustomerGateway#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The type of VPN connection that this customer gateway supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#type Ec2CustomerGateway#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// For devices that support BGP, the customer gateway's BGP ASN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#bgp_asn Ec2CustomerGateway#bgp_asn}
	BgpAsn *float64 `field:"optional" json:"bgpAsn" yaml:"bgpAsn"`
	// A name for the customer gateway device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#device_name Ec2CustomerGateway#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// One or more tags for the customer gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#tags Ec2CustomerGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

