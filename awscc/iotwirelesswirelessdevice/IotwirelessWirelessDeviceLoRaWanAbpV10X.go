package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanAbpV10X struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#dev_addr IotwirelessWirelessDevice#dev_addr}.
	DevAddr *string `field:"required" json:"devAddr" yaml:"devAddr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#session_keys IotwirelessWirelessDevice#session_keys}.
	SessionKeys *IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys `field:"required" json:"sessionKeys" yaml:"sessionKeys"`
}

