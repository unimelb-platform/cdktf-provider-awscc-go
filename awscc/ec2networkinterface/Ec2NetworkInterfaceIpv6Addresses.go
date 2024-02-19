package ec2networkinterface


type Ec2NetworkInterfaceIpv6Addresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#ipv_6_address Ec2NetworkInterface#ipv_6_address}.
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

