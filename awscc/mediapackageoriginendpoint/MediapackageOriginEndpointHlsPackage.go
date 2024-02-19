package mediapackageoriginendpoint


type MediapackageOriginEndpointHlsPackage struct {
	// This setting controls how ad markers are included in the packaged OriginEndpoint.
	//
	// "NONE" will omit all SCTE-35 ad markers from the output. "PASSTHROUGH" causes the manifest to contain a copy of the SCTE-35 ad markers (comments) taken directly from the input HTTP Live Streaming (HLS) manifest. "SCTE35_ENHANCED" generates ad markers and blackout tags based on SCTE-35 messages in the input source. "DATERANGE" inserts EXT-X-DATERANGE tags to signal ad and program transition events in HLS and CMAF manifests. For this option, you must set a programDateTimeIntervalSeconds value that is greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#ad_markers MediapackageOriginEndpoint#ad_markers}
	AdMarkers *string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// This setting allows the delivery restriction flags on SCTE-35 segmentation descriptors to determine whether a message signals an ad.
	//
	// Choosing "NONE" means no SCTE-35 messages become ads.  Choosing "RESTRICTED" means SCTE-35 messages of the types specified in AdTriggers that contain delivery restrictions will be treated as ads.  Choosing "UNRESTRICTED" means SCTE-35 messages of the types specified in AdTriggers that do not contain delivery restrictions will be treated as ads.  Choosing "BOTH" means all SCTE-35 messages of the types specified in AdTriggers will be treated as ads.  Note that Splice Insert messages do not have these flags and are always treated as ads if specified in AdTriggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#ads_on_delivery_restrictions MediapackageOriginEndpoint#ads_on_delivery_restrictions}
	AdsOnDeliveryRestrictions *string `field:"optional" json:"adsOnDeliveryRestrictions" yaml:"adsOnDeliveryRestrictions"`
	// A list of SCTE-35 message types that are treated as ad markers in the output.
	//
	// If empty, no ad markers are output.  Specify multiple items to create ad markers for all of the included message types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#ad_triggers MediapackageOriginEndpoint#ad_triggers}
	AdTriggers *[]*string `field:"optional" json:"adTriggers" yaml:"adTriggers"`
	// An HTTP Live Streaming (HLS) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption MediapackageOriginEndpoint#encryption}
	Encryption *MediapackageOriginEndpointHlsPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#include_dvb_subtitles MediapackageOriginEndpoint#include_dvb_subtitles}
	IncludeDvbSubtitles interface{} `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
	// When enabled, an I-Frame only stream will be included in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#include_iframe_only_stream MediapackageOriginEndpoint#include_iframe_only_stream}
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// The HTTP Live Streaming (HLS) playlist type.
	//
	// When either "EVENT" or "VOD" is specified, a corresponding EXT-X-PLAYLIST-TYPE entry will be included in the media playlist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#playlist_type MediapackageOriginEndpoint#playlist_type}
	PlaylistType *string `field:"optional" json:"playlistType" yaml:"playlistType"`
	// Time window (in seconds) contained in each parent manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#playlist_window_seconds MediapackageOriginEndpoint#playlist_window_seconds}
	PlaylistWindowSeconds *float64 `field:"optional" json:"playlistWindowSeconds" yaml:"playlistWindowSeconds"`
	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into manifests.
	//
	// Additionally, when an interval is specified ID3Timed Metadata messages will be generated every 5 seconds using the ingest time of the content. If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME tags will be inserted into manifests and no ID3Timed Metadata messages will be generated. Note that irrespective of this parameter, if any ID3 Timed Metadata is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#program_date_time_interval_seconds MediapackageOriginEndpoint#program_date_time_interval_seconds}
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// Duration (in seconds) of each fragment.
	//
	// Actual fragments will be rounded to the nearest multiple of the source fragment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#segment_duration_seconds MediapackageOriginEndpoint#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// A StreamSelection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#stream_selection MediapackageOriginEndpoint#stream_selection}
	StreamSelection *MediapackageOriginEndpointHlsPackageStreamSelection `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// When enabled, audio streams will be placed in rendition groups in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#use_audio_rendition_group MediapackageOriginEndpoint#use_audio_rendition_group}
	UseAudioRenditionGroup interface{} `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}

