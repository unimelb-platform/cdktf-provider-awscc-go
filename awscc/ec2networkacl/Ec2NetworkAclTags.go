package ec2networkacl


type Ec2NetworkAclTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_acl#key Ec2NetworkAcl#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_acl#value Ec2NetworkAcl#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

