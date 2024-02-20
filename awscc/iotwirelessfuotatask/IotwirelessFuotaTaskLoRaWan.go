package iotwirelessfuotatask


type IotwirelessFuotaTaskLoRaWan struct {
	// FUOTA task LoRaWAN RF region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_fuota_task#rf_region IotwirelessFuotaTask#rf_region}
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

