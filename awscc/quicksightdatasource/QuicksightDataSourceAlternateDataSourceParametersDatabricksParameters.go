package quicksightdatasource


type QuicksightDataSourceAlternateDataSourceParametersDatabricksParameters struct {
	// <p>Host.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// <p>Port.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// <p>The HTTP Path of the Databricks data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#sql_endpoint_path QuicksightDataSource#sql_endpoint_path}
	SqlEndpointPath *string `field:"required" json:"sqlEndpointPath" yaml:"sqlEndpointPath"`
}

