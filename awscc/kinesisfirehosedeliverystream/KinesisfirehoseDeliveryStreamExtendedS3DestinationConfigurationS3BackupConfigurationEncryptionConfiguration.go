package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#kms_encryption_config KinesisfirehoseDeliveryStream#kms_encryption_config}.
	KmsEncryptionConfig *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfigurationEncryptionConfigurationKmsEncryptionConfig `field:"optional" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#no_encryption_config KinesisfirehoseDeliveryStream#no_encryption_config}.
	NoEncryptionConfig *string `field:"optional" json:"noEncryptionConfig" yaml:"noEncryptionConfig"`
}

