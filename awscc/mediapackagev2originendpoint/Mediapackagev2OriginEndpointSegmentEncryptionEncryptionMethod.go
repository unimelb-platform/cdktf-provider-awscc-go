package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#cmaf_encryption_method Mediapackagev2OriginEndpoint#cmaf_encryption_method}.
	CmafEncryptionMethod *string `field:"optional" json:"cmafEncryptionMethod" yaml:"cmafEncryptionMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackagev2_origin_endpoint#ts_encryption_method Mediapackagev2OriginEndpoint#ts_encryption_method}.
	TsEncryptionMethod *string `field:"optional" json:"tsEncryptionMethod" yaml:"tsEncryptionMethod"`
}

