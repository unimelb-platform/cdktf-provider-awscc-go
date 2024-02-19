package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanAbpV10XSessionKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#app_s_key IotwirelessWirelessDevice#app_s_key}.
	AppSKey *string `field:"required" json:"appSKey" yaml:"appSKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#nwk_s_key IotwirelessWirelessDevice#nwk_s_key}.
	NwkSKey *string `field:"required" json:"nwkSKey" yaml:"nwkSKey"`
}

