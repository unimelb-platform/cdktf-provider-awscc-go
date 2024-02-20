package ec2prefixlist


type Ec2PrefixListEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#cidr Ec2PrefixList#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_prefix_list#description Ec2PrefixList#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

