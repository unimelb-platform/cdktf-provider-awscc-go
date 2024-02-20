package ec2verifiedaccessendpoint


type Ec2VerifiedAccessEndpointNetworkInterfaceOptions struct {
	// The ID of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint#network_interface_id Ec2VerifiedAccessEndpoint#network_interface_id}
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The IP port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint#port Ec2VerifiedAccessEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The IP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint#protocol Ec2VerifiedAccessEndpoint#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

