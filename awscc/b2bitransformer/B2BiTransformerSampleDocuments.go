package b2bitransformer


type B2BiTransformerSampleDocuments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#bucket_name B2BiTransformer#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_transformer#keys B2BiTransformer#keys}.
	Keys interface{} `field:"optional" json:"keys" yaml:"keys"`
}

