package ec2eip


type Ec2EipTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip#key Ec2Eip#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_eip#value Ec2Eip#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

