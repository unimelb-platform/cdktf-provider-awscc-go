package datazonedatasource


type DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#expression DatazoneDataSource#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The search filter expression type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#type DatazoneDataSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

