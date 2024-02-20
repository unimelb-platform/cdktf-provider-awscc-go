package ec2instance


type Ec2InstanceNetworkInterfacesPrivateIpAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#primary Ec2Instance#primary}.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#private_ip_address Ec2Instance#private_ip_address}.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
}

