package connectinstancestorageconfig


type ConnectInstanceStorageConfigS3Config struct {
	// A name for the S3 Bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#bucket_name ConnectInstanceStorageConfig#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Prefixes are used to infer logical hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#bucket_prefix ConnectInstanceStorageConfig#bucket_prefix}
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#encryption_config ConnectInstanceStorageConfig#encryption_config}.
	EncryptionConfig *ConnectInstanceStorageConfigS3ConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
}

