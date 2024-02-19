package ec2networkinterface


type Ec2NetworkInterfaceIpv4Prefixes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#ipv_4_prefix Ec2NetworkInterface#ipv_4_prefix}.
	Ipv4Prefix *string `field:"required" json:"ipv4Prefix" yaml:"ipv4Prefix"`
}

