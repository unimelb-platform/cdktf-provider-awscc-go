package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputSchemaRecordColumns struct {
	// The name of the column that is created in the in-application input stream or reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#name Kinesisanalyticsv2Application#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of column created in the in-application input stream or reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#sql_type Kinesisanalyticsv2Application#sql_type}
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// A reference to the data element in the streaming input or the reference data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#mapping Kinesisanalyticsv2Application#mapping}
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

