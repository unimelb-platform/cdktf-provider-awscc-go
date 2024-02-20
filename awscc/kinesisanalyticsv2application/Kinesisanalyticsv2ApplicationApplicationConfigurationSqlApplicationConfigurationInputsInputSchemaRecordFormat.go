package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordFormat struct {
	// The type of record format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#record_format_type Kinesisanalyticsv2Application#record_format_type}
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When you configure application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#mapping_parameters Kinesisanalyticsv2Application#mapping_parameters}
	MappingParameters *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordFormatMappingParameters `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

