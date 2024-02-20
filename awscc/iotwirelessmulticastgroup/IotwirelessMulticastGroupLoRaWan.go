package iotwirelessmulticastgroup


type IotwirelessMulticastGroupLoRaWan struct {
	// Multicast group LoRaWAN DL Class.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#dl_class IotwirelessMulticastGroup#dl_class}
	DlClass *string `field:"required" json:"dlClass" yaml:"dlClass"`
	// Multicast group LoRaWAN RF region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_multicast_group#rf_region IotwirelessMulticastGroup#rf_region}
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

