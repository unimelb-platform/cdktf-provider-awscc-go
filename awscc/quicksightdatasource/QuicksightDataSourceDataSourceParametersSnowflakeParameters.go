package quicksightdatasource


type QuicksightDataSourceDataSourceParametersSnowflakeParameters struct {
	// <p>Database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// <p>Host.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// <p>Warehouse.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#warehouse QuicksightDataSource#warehouse}
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
}

