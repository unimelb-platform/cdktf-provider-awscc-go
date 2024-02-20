package mediaconnectbridge


type MediaconnectBridgeEgressGatewayBridge struct {
	// The maximum expected bitrate of the egress bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#max_bitrate MediaconnectBridge#max_bitrate}
	MaxBitrate *float64 `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
}

