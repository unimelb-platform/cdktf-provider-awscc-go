package ec2ec2fleet


type Ec2Ec2FleetTagSpecificationsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#key Ec2Ec2Fleet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#value Ec2Ec2Fleet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

