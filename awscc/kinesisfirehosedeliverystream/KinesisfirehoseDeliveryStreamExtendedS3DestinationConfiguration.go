package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#bucket_arn KinesisfirehoseDeliveryStream#bucket_arn}.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#role_arn KinesisfirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#buffering_hints KinesisfirehoseDeliveryStream#buffering_hints}.
	BufferingHints *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationBufferingHints `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#cloudwatch_logging_options KinesisfirehoseDeliveryStream#cloudwatch_logging_options}.
	CloudwatchLoggingOptions *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationCloudwatchLoggingOptions `field:"optional" json:"cloudwatchLoggingOptions" yaml:"cloudwatchLoggingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#compression_format KinesisfirehoseDeliveryStream#compression_format}.
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#data_format_conversion_configuration KinesisfirehoseDeliveryStream#data_format_conversion_configuration}.
	DataFormatConversionConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfiguration `field:"optional" json:"dataFormatConversionConfiguration" yaml:"dataFormatConversionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#dynamic_partitioning_configuration KinesisfirehoseDeliveryStream#dynamic_partitioning_configuration}.
	DynamicPartitioningConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfiguration `field:"optional" json:"dynamicPartitioningConfiguration" yaml:"dynamicPartitioningConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#encryption_configuration KinesisfirehoseDeliveryStream#encryption_configuration}.
	EncryptionConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#error_output_prefix KinesisfirehoseDeliveryStream#error_output_prefix}.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#prefix KinesisfirehoseDeliveryStream#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#processing_configuration KinesisfirehoseDeliveryStream#processing_configuration}.
	ProcessingConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationProcessingConfiguration `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#s3_backup_configuration KinesisfirehoseDeliveryStream#s3_backup_configuration}.
	S3BackupConfiguration *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationS3BackupConfiguration `field:"optional" json:"s3BackupConfiguration" yaml:"s3BackupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#s3_backup_mode KinesisfirehoseDeliveryStream#s3_backup_mode}.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
}

