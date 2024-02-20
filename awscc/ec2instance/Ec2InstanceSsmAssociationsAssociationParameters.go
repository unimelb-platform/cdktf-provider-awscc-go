package ec2instance


type Ec2InstanceSsmAssociationsAssociationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#key Ec2Instance#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#value Ec2Instance#value}.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

