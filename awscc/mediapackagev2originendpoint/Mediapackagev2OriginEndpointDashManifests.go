package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointDashManifests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#drm_signaling Mediapackagev2OriginEndpoint#drm_signaling}.
	DrmSignaling *string `field:"optional" json:"drmSignaling" yaml:"drmSignaling"`
	// <p>Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#filter_configuration Mediapackagev2OriginEndpoint#filter_configuration}
	FilterConfiguration *Mediapackagev2OriginEndpointDashManifestsFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// <p>A short string that's appended to the endpoint URL.
	//
	// The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, index. </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#manifest_name Mediapackagev2OriginEndpoint#manifest_name}
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// <p>The total duration (in seconds) of the manifest's content.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#manifest_window_seconds Mediapackagev2OriginEndpoint#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// <p>Minimum amount of content (in seconds) that a player must keep available in the buffer.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#min_buffer_time_seconds Mediapackagev2OriginEndpoint#min_buffer_time_seconds}
	MinBufferTimeSeconds *float64 `field:"optional" json:"minBufferTimeSeconds" yaml:"minBufferTimeSeconds"`
	// <p>Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#min_update_period_seconds Mediapackagev2OriginEndpoint#min_update_period_seconds}
	MinUpdatePeriodSeconds *float64 `field:"optional" json:"minUpdatePeriodSeconds" yaml:"minUpdatePeriodSeconds"`
	// <p>A list of triggers that controls when AWS Elemental MediaPackage separates the MPEG-DASH manifest into multiple periods.
	//
	// Leave this value empty to indicate that the manifest is contained all in one period.
	//          For more information about periods in the DASH manifest, see <a href="https://docs.aws.amazon.com/mediapackage/latest/userguide/multi-period.html">Multi-period DASH in AWS Elemental MediaPackage</a>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#period_triggers Mediapackagev2OriginEndpoint#period_triggers}
	PeriodTriggers *[]*string `field:"optional" json:"periodTriggers" yaml:"periodTriggers"`
	// <p>The SCTE configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#scte_dash Mediapackagev2OriginEndpoint#scte_dash}
	ScteDash *Mediapackagev2OriginEndpointDashManifestsScteDash `field:"optional" json:"scteDash" yaml:"scteDash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#segment_template_format Mediapackagev2OriginEndpoint#segment_template_format}.
	SegmentTemplateFormat *string `field:"optional" json:"segmentTemplateFormat" yaml:"segmentTemplateFormat"`
	// <p>The amount of time (in seconds) that the player should be from the end of the manifest.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#suggested_presentation_delay_seconds Mediapackagev2OriginEndpoint#suggested_presentation_delay_seconds}
	SuggestedPresentationDelaySeconds *float64 `field:"optional" json:"suggestedPresentationDelaySeconds" yaml:"suggestedPresentationDelaySeconds"`
	// <p>Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_origin_endpoint#utc_timing Mediapackagev2OriginEndpoint#utc_timing}
	UtcTiming *Mediapackagev2OriginEndpointDashManifestsUtcTiming `field:"optional" json:"utcTiming" yaml:"utcTiming"`
}

