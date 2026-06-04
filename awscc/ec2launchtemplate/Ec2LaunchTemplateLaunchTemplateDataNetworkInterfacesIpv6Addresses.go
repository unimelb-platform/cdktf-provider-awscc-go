package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6Addresses struct {
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet.
	//
	// You can't use this option if you're specifying a number of IPv6 addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_launch_template#ipv_6_address Ec2LaunchTemplate#ipv_6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
}

