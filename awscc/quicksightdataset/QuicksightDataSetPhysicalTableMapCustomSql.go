package quicksightdataset


type QuicksightDataSetPhysicalTableMapCustomSql struct {
	// <p>The column schema from the SQL query result set.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// <p>The Amazon Resource Name (ARN) of the data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// <p>A display name for the SQL query result.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The SQL query.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#sql_query QuicksightDataSet#sql_query}
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
}

