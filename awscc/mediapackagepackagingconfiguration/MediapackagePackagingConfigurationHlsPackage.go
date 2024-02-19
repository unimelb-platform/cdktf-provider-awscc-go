package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationHlsPackage struct {
	// A list of HLS manifest configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#hls_manifests MediapackagePackagingConfiguration#hls_manifests}
	HlsManifests interface{} `field:"required" json:"hlsManifests" yaml:"hlsManifests"`
	// An HTTP Live Streaming (HLS) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#encryption MediapackagePackagingConfiguration#encryption}
	Encryption *MediapackagePackagingConfigurationHlsPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#include_dvb_subtitles MediapackagePackagingConfiguration#include_dvb_subtitles}
	IncludeDvbSubtitles interface{} `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments will be rounded to the nearest multiple of the source fragment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#segment_duration_seconds MediapackagePackagingConfiguration#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// When enabled, audio streams will be placed in rendition groups in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#use_audio_rendition_group MediapackagePackagingConfiguration#use_audio_rendition_group}
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

