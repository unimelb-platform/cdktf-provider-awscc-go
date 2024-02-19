package mediaconnectbridgesource


type MediaconnectBridgeSourceNetworkSource struct {
	// The network source multicast IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#multicast_ip MediaconnectBridgeSource#multicast_ip}
	MulticastIp *string `field:"required" json:"multicastIp" yaml:"multicastIp"`
	// The network source's gateway network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#network_name MediaconnectBridgeSource#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network source port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#port MediaconnectBridgeSource#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network source protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#protocol MediaconnectBridgeSource#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

