package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointLowLatencyHlsManifestsFilterConfiguration struct {
	// <p>Optionally specify the end time for all of your manifest egress requests.
	//
	// When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#end Mediapackagev2OriginEndpoint#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// <p>Optionally specify one or more manifest filters for all of your manifest egress requests.
	//
	// When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#manifest_filter Mediapackagev2OriginEndpoint#manifest_filter}
	ManifestFilter *string `field:"optional" json:"manifestFilter" yaml:"manifestFilter"`
	// <p>Optionally specify the start time for all of your manifest egress requests.
	//
	// When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#start Mediapackagev2OriginEndpoint#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
	// <p>Optionally specify the time delay for all of your manifest egress requests.
	//
	// Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#time_delay_seconds Mediapackagev2OriginEndpoint#time_delay_seconds}
	TimeDelaySeconds *float64 `field:"optional" json:"timeDelaySeconds" yaml:"timeDelaySeconds"`
}

