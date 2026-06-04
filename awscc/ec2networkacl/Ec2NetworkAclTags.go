package ec2networkacl


type Ec2NetworkAclTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_network_acl#key Ec2NetworkAcl#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_network_acl#value Ec2NetworkAcl#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

