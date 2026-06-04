package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationDestinationTableConfigurationListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#destination_database_name KinesisfirehoseDeliveryStream#destination_database_name}.
	DestinationDatabaseName *string `field:"optional" json:"destinationDatabaseName" yaml:"destinationDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#destination_table_name KinesisfirehoseDeliveryStream#destination_table_name}.
	DestinationTableName *string `field:"optional" json:"destinationTableName" yaml:"destinationTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#s3_error_output_prefix KinesisfirehoseDeliveryStream#s3_error_output_prefix}.
	S3ErrorOutputPrefix *string `field:"optional" json:"s3ErrorOutputPrefix" yaml:"s3ErrorOutputPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#unique_keys KinesisfirehoseDeliveryStream#unique_keys}.
	UniqueKeys *[]*string `field:"optional" json:"uniqueKeys" yaml:"uniqueKeys"`
}

