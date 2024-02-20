package s3bucket


type S3BucketReplicationConfigurationRulesDestinationEncryptionConfiguration struct {
	// Specifies the ID (Key ARN or Alias ARN) of the customer managed customer master key (CMK) stored in AWS Key Management Service (KMS) for the destination bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#replica_kms_key_id S3Bucket#replica_kms_key_id}
	ReplicaKmsKeyId *string `field:"required" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
}

