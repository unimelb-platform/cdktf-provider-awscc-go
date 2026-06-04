package mediaconnectflow


type MediaconnectFlowSourceFailoverConfigSourcePriority struct {
	// The name of the source you choose as the primary source for this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediaconnect_flow#primary_source MediaconnectFlow#primary_source}
	PrimarySource *string `field:"optional" json:"primarySource" yaml:"primarySource"`
}

