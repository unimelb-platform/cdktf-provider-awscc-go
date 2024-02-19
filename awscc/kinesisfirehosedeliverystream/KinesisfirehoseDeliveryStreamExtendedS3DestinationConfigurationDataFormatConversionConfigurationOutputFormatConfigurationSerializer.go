package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#orc_ser_de KinesisfirehoseDeliveryStream#orc_ser_de}.
	OrcSerDe *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe `field:"optional" json:"orcSerDe" yaml:"orcSerDe"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#parquet_ser_de KinesisfirehoseDeliveryStream#parquet_ser_de}.
	ParquetSerDe *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe `field:"optional" json:"parquetSerDe" yaml:"parquetSerDe"`
}

