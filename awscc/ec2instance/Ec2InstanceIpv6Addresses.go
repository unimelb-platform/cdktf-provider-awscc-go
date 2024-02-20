package ec2instance


type Ec2InstanceIpv6Addresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#ipv_6_address Ec2Instance#ipv_6_address}.
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

