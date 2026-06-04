package ec2networkinterface


type Ec2NetworkInterfaceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_network_interface#key Ec2NetworkInterface#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_network_interface#value Ec2NetworkInterface#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

