package iotwirelessdeviceprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotwirelessDeviceProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_device_profile#lo_ra_wan IotwirelessDeviceProfile#lo_ra_wan}
	LoRaWan *IotwirelessDeviceProfileLoRaWan `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// Name of service profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_device_profile#name IotwirelessDeviceProfile#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that contain metadata for the device profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_device_profile#tags IotwirelessDeviceProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

