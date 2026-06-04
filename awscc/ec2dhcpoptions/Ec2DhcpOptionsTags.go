package ec2dhcpoptions


type Ec2DhcpOptionsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_dhcp_options#key Ec2DhcpOptions#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_dhcp_options#value Ec2DhcpOptions#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

