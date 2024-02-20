package quicksightdatasource


type QuicksightDataSourceDataSourceParametersSparkParameters struct {
	// <p>Host.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// <p>Port.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

