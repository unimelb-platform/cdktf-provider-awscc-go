package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanOtaaV10X struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#app_eui IotwirelessWirelessDevice#app_eui}.
	AppEui *string `field:"required" json:"appEui" yaml:"appEui"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#app_key IotwirelessWirelessDevice#app_key}.
	AppKey *string `field:"required" json:"appKey" yaml:"appKey"`
}

