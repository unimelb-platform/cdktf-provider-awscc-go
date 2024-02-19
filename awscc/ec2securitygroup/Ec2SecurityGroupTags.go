package ec2securitygroup


type Ec2SecurityGroupTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group#key Ec2SecurityGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group#value Ec2SecurityGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

