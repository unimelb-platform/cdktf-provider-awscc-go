package networkmanagerdevice


type NetworkmanagerDeviceLocation struct {
	// The physical address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_device#address NetworkmanagerDevice#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The latitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_device#latitude NetworkmanagerDevice#latitude}
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// The longitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_device#longitude NetworkmanagerDevice#longitude}
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

