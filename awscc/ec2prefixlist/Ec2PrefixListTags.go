package ec2prefixlist


type Ec2PrefixListTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#key Ec2PrefixList#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#value Ec2PrefixList#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

