package mediaconnectbridge


type MediaconnectBridgeSourceFailoverConfigSourcePriority struct {
	// The name of the source you choose as the primary source for this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#primary_source MediaconnectBridge#primary_source}
	PrimarySource *string `field:"optional" json:"primarySource" yaml:"primarySource"`
}

