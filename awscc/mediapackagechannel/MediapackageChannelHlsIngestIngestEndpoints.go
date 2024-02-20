package mediapackagechannel


type MediapackageChannelHlsIngestIngestEndpoints struct {
	// The system generated unique identifier for the IngestEndpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#id MediapackageChannel#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The system generated password for ingest authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#password MediapackageChannel#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The ingest URL to which the source stream should be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#url MediapackageChannel#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The system generated username for ingest authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_channel#username MediapackageChannel#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

