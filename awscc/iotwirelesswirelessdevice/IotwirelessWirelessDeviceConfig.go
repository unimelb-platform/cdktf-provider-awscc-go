package iotwirelesswirelessdevice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotwirelessWirelessDeviceConfig struct {
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
	// Wireless device destination name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#destination_name IotwirelessWirelessDevice#destination_name}
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// Wireless device type, currently only Sidewalk and LoRa.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#type IotwirelessWirelessDevice#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Wireless device description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#description IotwirelessWirelessDevice#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time when the most recent uplink was received.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#last_uplink_received_at IotwirelessWirelessDevice#last_uplink_received_at}
	LastUplinkReceivedAt *string `field:"optional" json:"lastUplinkReceivedAt" yaml:"lastUplinkReceivedAt"`
	// The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#lo_ra_wan IotwirelessWirelessDevice#lo_ra_wan}
	LoRaWan *IotwirelessWirelessDeviceLoRaWan `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// Wireless device name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#name IotwirelessWirelessDevice#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of key-value pairs that contain metadata for the device.
	//
	// Currently not supported, will not create if tags are passed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#tags IotwirelessWirelessDevice#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Thing arn. Passed into update to associate Thing with Wireless device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#thing_arn IotwirelessWirelessDevice#thing_arn}
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
}

