package ec2instance


type Ec2InstanceIpv6Addresses struct {
	// The IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_instance#ipv_6_address Ec2Instance#ipv_6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
}

