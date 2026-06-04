package quicksightdatasource


type QuicksightDataSourceDataSourceParametersOracleParameters struct {
	// <p>The database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// <p>An Oracle host.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// <p>The port.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

