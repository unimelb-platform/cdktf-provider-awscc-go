package mediaconnectbridge


type MediaconnectBridgeOutputsNetworkOutput struct {
	// The network output IP Address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#ip_address MediaconnectBridge#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The network output name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network output's gateway network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#network_name MediaconnectBridge#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network output port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#port MediaconnectBridge#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network output protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#protocol MediaconnectBridge#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The network output TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#ttl MediaconnectBridge#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

