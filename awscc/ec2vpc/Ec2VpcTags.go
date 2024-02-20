package ec2vpc


type Ec2VpcTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc#key Ec2Vpc#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc#value Ec2Vpc#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

