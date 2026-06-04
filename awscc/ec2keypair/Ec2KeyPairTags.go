package ec2keypair


type Ec2KeyPairTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#key Ec2KeyPair#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_key_pair#value Ec2KeyPair#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

