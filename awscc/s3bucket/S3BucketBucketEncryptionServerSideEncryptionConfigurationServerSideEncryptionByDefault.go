package s3bucket


type S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#sse_algorithm S3Bucket#sse_algorithm}.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// "KMSMasterKeyID" can only be used when you set the value of SSEAlgorithm as aws:kms or aws:kms:dsse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#kms_master_key_id S3Bucket#kms_master_key_id}
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

