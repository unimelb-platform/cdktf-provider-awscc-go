package quicksightdataset


type QuicksightDataSetPhysicalTableMapRelationalTable struct {
	// <p>The Amazon Resource Name (ARN) for the data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// <p>The column schema of the table.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#input_columns QuicksightDataSet#input_columns}
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// <p>The name of the relational table.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The catalog associated with a table.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#catalog QuicksightDataSet#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// <p>The schema name. This name applies to certain relational database engines.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#schema QuicksightDataSet#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

