package ec2networkinterface


type Ec2NetworkInterfaceIpv6Prefixes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#ipv_6_prefix Ec2NetworkInterface#ipv_6_prefix}.
	Ipv6Prefix *string `field:"required" json:"ipv6Prefix" yaml:"ipv6Prefix"`
}

