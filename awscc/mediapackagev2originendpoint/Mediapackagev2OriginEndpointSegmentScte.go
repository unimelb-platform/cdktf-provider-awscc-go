package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentScte struct {
	// <p>The SCTE-35 message types that you want to be treated as ad markers in the output.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#scte_filter Mediapackagev2OriginEndpoint#scte_filter}
	ScteFilter *[]*string `field:"optional" json:"scteFilter" yaml:"scteFilter"`
}

