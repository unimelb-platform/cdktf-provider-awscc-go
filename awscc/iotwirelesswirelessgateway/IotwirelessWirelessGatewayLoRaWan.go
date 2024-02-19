package iotwirelesswirelessgateway


type IotwirelessWirelessGatewayLoRaWan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_gateway#gateway_eui IotwirelessWirelessGateway#gateway_eui}.
	GatewayEui *string `field:"required" json:"gatewayEui" yaml:"gatewayEui"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_wireless_gateway#rf_region IotwirelessWirelessGateway#rf_region}.
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
}

