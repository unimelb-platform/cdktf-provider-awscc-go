package connectinstancestorageconfig


type ConnectInstanceStorageConfigKinesisVideoStreamConfigEncryptionConfig struct {
	// Specifies default encryption using AWS KMS-Managed Keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#encryption_type ConnectInstanceStorageConfig#encryption_type}
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// Specifies the encryption key id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#key_id ConnectInstanceStorageConfig#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

