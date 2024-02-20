package mediapackageoriginendpoint


type MediapackageOriginEndpointHlsPackageStreamSelection struct {
	// The maximum video bitrate (bps) to include in output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#max_video_bits_per_second MediapackageOriginEndpoint#max_video_bits_per_second}
	MaxVideoBitsPerSecond *float64 `field:"optional" json:"maxVideoBitsPerSecond" yaml:"maxVideoBitsPerSecond"`
	// The minimum video bitrate (bps) to include in output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#min_video_bits_per_second MediapackageOriginEndpoint#min_video_bits_per_second}
	MinVideoBitsPerSecond *float64 `field:"optional" json:"minVideoBitsPerSecond" yaml:"minVideoBitsPerSecond"`
	// A directive that determines the order of streams in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#stream_order MediapackageOriginEndpoint#stream_order}
	StreamOrder *string `field:"optional" json:"streamOrder" yaml:"streamOrder"`
}

