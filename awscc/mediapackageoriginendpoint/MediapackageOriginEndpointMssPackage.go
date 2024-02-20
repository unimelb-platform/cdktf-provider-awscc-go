package mediapackageoriginendpoint


type MediapackageOriginEndpointMssPackage struct {
	// A Microsoft Smooth Streaming (MSS) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption MediapackageOriginEndpoint#encryption}
	Encryption *MediapackageOriginEndpointMssPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The time window (in seconds) contained in each manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#manifest_window_seconds MediapackageOriginEndpoint#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// The duration (in seconds) of each segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#segment_duration_seconds MediapackageOriginEndpoint#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// A StreamSelection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#stream_selection MediapackageOriginEndpoint#stream_selection}
	StreamSelection *MediapackageOriginEndpointMssPackageStreamSelection `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

