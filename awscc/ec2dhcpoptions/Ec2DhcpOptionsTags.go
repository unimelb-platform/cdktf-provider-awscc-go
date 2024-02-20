package ec2dhcpoptions


type Ec2DhcpOptionsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#key Ec2DhcpOptions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#value Ec2DhcpOptions#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

