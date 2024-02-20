package mediapackageoriginendpoint


type MediapackageOriginEndpointCmafPackage struct {
	// A Common Media Application Format (CMAF) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption MediapackageOriginEndpoint#encryption}
	Encryption *MediapackageOriginEndpointCmafPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#hls_manifests MediapackageOriginEndpoint#hls_manifests}
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments will be rounded to the nearest multiple of the source segment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#segment_duration_seconds MediapackageOriginEndpoint#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// An optional custom string that is prepended to the name of each segment.
	//
	// If not specified, it defaults to the ChannelId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#segment_prefix MediapackageOriginEndpoint#segment_prefix}
	SegmentPrefix *string `field:"optional" json:"segmentPrefix" yaml:"segmentPrefix"`
	// A StreamSelection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#stream_selection MediapackageOriginEndpoint#stream_selection}
	StreamSelection *MediapackageOriginEndpointCmafPackageStreamSelection `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}

