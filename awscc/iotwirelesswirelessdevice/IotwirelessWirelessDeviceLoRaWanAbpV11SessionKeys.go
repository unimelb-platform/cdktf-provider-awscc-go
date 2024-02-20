package iotwirelesswirelessdevice


type IotwirelessWirelessDeviceLoRaWanAbpV11SessionKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#app_s_key IotwirelessWirelessDevice#app_s_key}.
	AppSKey *string `field:"required" json:"appSKey" yaml:"appSKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#f_nwk_s_int_key IotwirelessWirelessDevice#f_nwk_s_int_key}.
	FNwkSIntKey *string `field:"required" json:"fNwkSIntKey" yaml:"fNwkSIntKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#nwk_s_enc_key IotwirelessWirelessDevice#nwk_s_enc_key}.
	NwkSEncKey *string `field:"required" json:"nwkSEncKey" yaml:"nwkSEncKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_device#s_nwk_s_int_key IotwirelessWirelessDevice#s_nwk_s_int_key}.
	SNwkSIntKey *string `field:"required" json:"sNwkSIntKey" yaml:"sNwkSIntKey"`
}

