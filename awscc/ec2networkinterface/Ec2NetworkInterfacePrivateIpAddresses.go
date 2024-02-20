package ec2networkinterface


type Ec2NetworkInterfacePrivateIpAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#primary Ec2NetworkInterface#primary}.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#private_ip_address Ec2NetworkInterface#private_ip_address}.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
}

