package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationS3ConfigurationBufferingHints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#interval_in_seconds KinesisfirehoseDeliveryStream#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#size_in_m_bs KinesisfirehoseDeliveryStream#size_in_m_bs}.
	SizeInMBs *float64 `field:"optional" json:"sizeInMBs" yaml:"sizeInMBs"`
}

