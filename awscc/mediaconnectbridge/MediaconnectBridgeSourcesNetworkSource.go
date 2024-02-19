package mediaconnectbridge


type MediaconnectBridgeSourcesNetworkSource struct {
	// The network source multicast IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#multicast_ip MediaconnectBridge#multicast_ip}
	MulticastIp *string `field:"required" json:"multicastIp" yaml:"multicastIp"`
	// The name of the network source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network source's gateway network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#network_name MediaconnectBridge#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network source port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#port MediaconnectBridge#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network source protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#protocol MediaconnectBridge#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

