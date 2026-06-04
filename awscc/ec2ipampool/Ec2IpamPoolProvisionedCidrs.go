package ec2ipampool


type Ec2IpamPoolProvisionedCidrs struct {
	// Represents a single IPv4 or IPv6 CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_ipam_pool#cidr Ec2IpamPool#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

