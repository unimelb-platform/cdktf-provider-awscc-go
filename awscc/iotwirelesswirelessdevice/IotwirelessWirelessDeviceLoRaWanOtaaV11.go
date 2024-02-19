package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanOtaaV11 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#app_key IotwirelessWirelessDevice#app_key}.
	AppKey *string `field:"required" json:"appKey" yaml:"appKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#join_eui IotwirelessWirelessDevice#join_eui}.
	JoinEui *string `field:"required" json:"joinEui" yaml:"joinEui"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#nwk_key IotwirelessWirelessDevice#nwk_key}.
	NwkKey *string `field:"required" json:"nwkKey" yaml:"nwkKey"`
}

