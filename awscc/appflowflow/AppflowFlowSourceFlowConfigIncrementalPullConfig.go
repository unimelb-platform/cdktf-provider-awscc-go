package appflowflow


type AppflowFlowSourceFlowConfigIncrementalPullConfig struct {
	// Name of the datetime/timestamp data type field to be used for importing incremental records from the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#datetime_type_field_name AppflowFlow#datetime_type_field_name}
	DatetimeTypeFieldName *string `field:"optional" json:"datetimeTypeFieldName" yaml:"datetimeTypeFieldName"`
}

