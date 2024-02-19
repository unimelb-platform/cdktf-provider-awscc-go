package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationCopyCommand struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#data_table_name KinesisfirehoseDeliveryStream#data_table_name}.
	DataTableName *string `field:"required" json:"dataTableName" yaml:"dataTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#copy_options KinesisfirehoseDeliveryStream#copy_options}.
	CopyOptions *string `field:"optional" json:"copyOptions" yaml:"copyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#data_table_columns KinesisfirehoseDeliveryStream#data_table_columns}.
	DataTableColumns *string `field:"optional" json:"dataTableColumns" yaml:"dataTableColumns"`
}

