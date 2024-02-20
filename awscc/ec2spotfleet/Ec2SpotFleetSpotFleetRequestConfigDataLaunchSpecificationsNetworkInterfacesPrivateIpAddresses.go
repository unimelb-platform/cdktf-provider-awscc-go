package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsNetworkInterfacesPrivateIpAddresses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_spot_fleet#private_ip_address Ec2SpotFleet#private_ip_address}.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_spot_fleet#primary Ec2SpotFleet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

