package ec2networkinterface


type Ec2NetworkInterfaceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#key Ec2NetworkInterface#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#value Ec2NetworkInterface#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

