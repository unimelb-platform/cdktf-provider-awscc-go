package s3bucket


type S3BucketReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	// Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_bucket#status S3Bucket#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

