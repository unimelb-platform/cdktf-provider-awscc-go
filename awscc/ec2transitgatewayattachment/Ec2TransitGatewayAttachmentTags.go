package ec2transitgatewayattachment


type Ec2TransitGatewayAttachmentTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#key Ec2TransitGatewayAttachment#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#value Ec2TransitGatewayAttachment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

