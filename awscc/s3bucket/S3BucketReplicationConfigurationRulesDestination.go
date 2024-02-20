package s3bucket


type S3BucketReplicationConfigurationRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#bucket S3Bucket#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket.
	//
	// If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#access_control_translation S3Bucket#access_control_translation}
	AccessControlTranslation *S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#account S3Bucket#account}.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#encryption_configuration S3Bucket#encryption_configuration}
	EncryptionConfiguration *S3BucketReplicationConfigurationRulesDestinationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#metrics S3Bucket#metrics}.
	Metrics *S3BucketReplicationConfigurationRulesDestinationMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#replication_time S3Bucket#replication_time}.
	ReplicationTime *S3BucketReplicationConfigurationRulesDestinationReplicationTime `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#storage_class S3Bucket#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

