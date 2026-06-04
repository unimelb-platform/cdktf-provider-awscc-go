package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4Prefixes struct {
	// The IPv4 prefix. For information, see [Assigning prefixes to network interfaces](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the *Amazon EC2 User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_launch_template#ipv_4_prefix Ec2LaunchTemplate#ipv_4_prefix}
	Ipv4Prefix *string `field:"optional" json:"ipv4Prefix" yaml:"ipv4Prefix"`
}

