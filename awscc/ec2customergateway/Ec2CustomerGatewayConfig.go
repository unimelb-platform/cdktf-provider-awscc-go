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
	// IPv4 address for the customer gateway device's outside interface.
	//
	// The address must be static. If ``OutsideIpAddressType`` in your VPN connection options is set to ``PrivateIpv4``, you can use an RFC6598 or RFC1918 private IPv4 address. If ``OutsideIpAddressType`` is set to ``PublicIpv4``, you can use a public IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#ip_address Ec2CustomerGateway#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The type of VPN connection that this customer gateway supports (``ipsec.1``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#type Ec2CustomerGateway#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// For customer gateway devices that support BGP, specify the device's ASN.
	//
	// You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``.
	//  Default: 65000
	//  Valid values: ``1`` to ``2,147,483,647``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#bgp_asn Ec2CustomerGateway#bgp_asn}
	BgpAsn *float64 `field:"optional" json:"bgpAsn" yaml:"bgpAsn"`
	// For customer gateway devices that support BGP, specify the device's ASN.
	//
	// You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``.
	//  Valid values: ``2,147,483,648`` to ``4,294,967,295``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#bgp_asn_extended Ec2CustomerGateway#bgp_asn_extended}
	BgpAsnExtended *float64 `field:"optional" json:"bgpAsnExtended" yaml:"bgpAsnExtended"`
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#certificate_arn Ec2CustomerGateway#certificate_arn}
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The name of customer gateway device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#device_name Ec2CustomerGateway#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// One or more tags for the customer gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#tags Ec2CustomerGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

