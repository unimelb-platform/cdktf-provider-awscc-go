package ec2subnet


type Ec2SubnetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#key Ec2Subnet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#value Ec2Subnet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

