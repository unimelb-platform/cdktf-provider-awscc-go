package ec2host


type Ec2HostTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_host#key Ec2Host#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_host#value Ec2Host#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

