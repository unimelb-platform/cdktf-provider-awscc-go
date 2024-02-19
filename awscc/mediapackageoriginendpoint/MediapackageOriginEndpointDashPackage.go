package mediapackageoriginendpoint


type MediapackageOriginEndpointDashPackage struct {
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
	// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption MediapackageOriginEndpoint#encryption}
	Encryption *MediapackageOriginEndpointDashPackageEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// When enabled, an I-Frame only stream will be included in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#include_iframe_only_stream MediapackageOriginEndpoint#include_iframe_only_stream}
	IncludeIframeOnlyStream interface{} `field:"optional" json:"includeIframeOnlyStream" yaml:"includeIframeOnlyStream"`
	// Determines the position of some tags in the Media Presentation Description (MPD).
	//
	// When set to FULL, elements like SegmentTemplate and ContentProtection are included in each Representation.  When set to COMPACT, duplicate elements are combined and presented at the AdaptationSet level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#manifest_layout MediapackageOriginEndpoint#manifest_layout}
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// Time window (in seconds) contained in each manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#manifest_window_seconds MediapackageOriginEndpoint#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Minimum duration (in seconds) that a player will buffer media before starting the presentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#min_buffer_time_seconds MediapackageOriginEndpoint#min_buffer_time_seconds}
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// Minimum duration (in seconds) between potential changes to the Dynamic Adaptive Streaming over HTTP (DASH) Media Presentation Description (MPD).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#min_update_period_seconds MediapackageOriginEndpoint#min_update_period_seconds}
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// A list of triggers that controls when the outgoing Dynamic Adaptive Streaming over HTTP (DASH) Media Presentation Description (MPD) will be partitioned into multiple periods.
	//
	// If empty, the content will not be partitioned into more than one period. If the list contains "ADS", new periods will be created where the Channel source contains SCTE-35 ad markers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#period_triggers MediapackageOriginEndpoint#period_triggers}
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// The Dynamic Adaptive Streaming over HTTP (DASH) profile type.
	//
	// When set to "HBBTV_1_5", HbbTV 1.5 compliant output is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#profile MediapackageOriginEndpoint#profile}
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments will be rounded to the nearest multiple of the source segment duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#segment_duration_seconds MediapackageOriginEndpoint#segment_duration_seconds}
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// Determines the type of SegmentTemplate included in the Media Presentation Description (MPD).
	//
	// When set to NUMBER_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Number$ media URLs.  When set to TIME_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Time$ media URLs. When set to NUMBER_WITH_DURATION, only a duration is included in each SegmentTemplate, with $Number$ media URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#segment_template_format MediapackageOriginEndpoint#segment_template_format}
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// A StreamSelection configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#stream_selection MediapackageOriginEndpoint#stream_selection}
	StreamSelection *MediapackageOriginEndpointDashPackageStreamSelection `field:"optional" json:"streamSelection" yaml:"streamSelection"`
	// Duration (in seconds) to delay live content before presentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#suggested_presentation_delay_seconds MediapackageOriginEndpoint#suggested_presentation_delay_seconds}
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
	// Determines the type of UTCTiming included in the Media Presentation Description (MPD).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#utc_timing MediapackageOriginEndpoint#utc_timing}
	UtcTiming *string `field:"optional" json:"utcTiming" yaml:"utcTiming"`
	// Specifies the value attribute of the UTCTiming field when utcTiming is set to HTTP-ISO, HTTP-HEAD or HTTP-XSDATE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#utc_timing_uri MediapackageOriginEndpoint#utc_timing_uri}
	UtcTimingUri *string `field:"optional" json:"utcTimingUri" yaml:"utcTimingUri"`
}

