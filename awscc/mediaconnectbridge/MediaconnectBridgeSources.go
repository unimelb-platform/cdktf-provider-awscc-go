package mediaconnectbridge


type MediaconnectBridgeSources struct {
	// The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#flow_source MediaconnectBridge#flow_source}
	FlowSource *MediaconnectBridgeSourcesFlowSource `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the bridge. A network source originates at your premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#network_source MediaconnectBridge#network_source}
	NetworkSource *MediaconnectBridgeSourcesNetworkSource `field:"optional" json:"networkSource" yaml:"networkSource"`
}

