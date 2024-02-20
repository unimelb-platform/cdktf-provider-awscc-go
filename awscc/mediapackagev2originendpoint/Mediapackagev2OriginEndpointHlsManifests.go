package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointHlsManifests struct {
	// <p>A short short string that's appended to the endpoint URL.
	//
	// The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, index. MediaPackage automatically inserts the format extension, such as .m3u8. You can't use the same manifest name if you use HLS manifest and low-latency HLS manifest. The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#manifest_name Mediapackagev2OriginEndpoint#manifest_name}
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// <p>A short string that's appended to the endpoint URL.
	//
	// The child manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default child manifest name, index_1. The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#child_manifest_name Mediapackagev2OriginEndpoint#child_manifest_name}
	ChildManifestName *string `field:"optional" json:"childManifestName" yaml:"childManifestName"`
	// <p>Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#filter_configuration Mediapackagev2OriginEndpoint#filter_configuration}
	FilterConfiguration *Mediapackagev2OriginEndpointHlsManifestsFilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// <p>The total duration (in seconds) of the manifest's content.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#manifest_window_seconds Mediapackagev2OriginEndpoint#manifest_window_seconds}
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// <p>Inserts EXT-X-PROGRAM-DATE-TIME tags in the output manifest at the interval that you specify.
	//
	// If you don't enter an interval,
	//          EXT-X-PROGRAM-DATE-TIME tags aren't included in the manifest.
	//          The tags sync the stream to the wall clock so that viewers can seek to a specific time in the playback timeline on the player.
	//          ID3Timed metadata messages generate every 5 seconds whenever the content is ingested.</p>
	//          <p>Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#program_date_time_interval_seconds Mediapackagev2OriginEndpoint#program_date_time_interval_seconds}
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// <p>The SCTE configuration.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#scte_hls Mediapackagev2OriginEndpoint#scte_hls}
	ScteHls *Mediapackagev2OriginEndpointHlsManifestsScteHls `field:"optional" json:"scteHls" yaml:"scteHls"`
	// <p>The egress domain URL for stream delivery from MediaPackage.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#url Mediapackagev2OriginEndpoint#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

