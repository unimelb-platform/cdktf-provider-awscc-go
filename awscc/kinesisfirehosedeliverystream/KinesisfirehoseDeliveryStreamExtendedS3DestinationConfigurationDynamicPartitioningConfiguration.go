package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#enabled KinesisfirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#retry_options KinesisfirehoseDeliveryStream#retry_options}.
	RetryOptions *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDynamicPartitioningConfigurationRetryOptions `field:"optional" json:"retryOptions" yaml:"retryOptions"`
}

