package ec2transitgatewayvpcattachment


type Ec2TransitGatewayVpcAttachmentTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_vpc_attachment#key Ec2TransitGatewayVpcAttachment#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_vpc_attachment#value Ec2TransitGatewayVpcAttachment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

