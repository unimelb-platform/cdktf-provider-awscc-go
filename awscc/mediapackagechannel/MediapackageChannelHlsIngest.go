package mediapackagechannel


type MediapackageChannelHlsIngest struct {
	// A list of endpoints to which the source stream should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#ingest_endpoints MediapackageChannel#ingest_endpoints}
	IngestEndpoints interface{} `field:"optional" json:"ingestEndpoints" yaml:"ingestEndpoints"`
}

