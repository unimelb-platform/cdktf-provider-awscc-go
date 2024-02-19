package ec2transitgatewayvpcattachment


type Ec2TransitGatewayVpcAttachmentOptions struct {
	// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_vpc_attachment#appliance_mode_support Ec2TransitGatewayVpcAttachment#appliance_mode_support}
	ApplianceModeSupport *string `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_vpc_attachment#dns_support Ec2TransitGatewayVpcAttachment#dns_support}
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_vpc_attachment#ipv_6_support Ec2TransitGatewayVpcAttachment#ipv_6_support}
	Ipv6Support *string `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

