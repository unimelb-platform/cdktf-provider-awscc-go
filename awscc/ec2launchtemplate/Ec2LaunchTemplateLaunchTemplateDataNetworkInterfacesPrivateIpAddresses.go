package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddresses struct {
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#primary Ec2LaunchTemplate#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The private IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#private_ip_address Ec2LaunchTemplate#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

