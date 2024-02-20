package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordFormatMappingParametersJsonMappingParameters struct {
	// The path to the top-level parent that contains the records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#record_row_path Kinesisanalyticsv2Application#record_row_path}
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

