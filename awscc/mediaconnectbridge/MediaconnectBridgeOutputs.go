package mediaconnectbridge


type MediaconnectBridgeOutputs struct {
	// The output of the bridge. A network output is delivered to your premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#network_output MediaconnectBridge#network_output}
	NetworkOutput *MediaconnectBridgeOutputsNetworkOutput `field:"optional" json:"networkOutput" yaml:"networkOutput"`
}

