package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperationColumns struct {
	// <p>A unique ID to identify a calculated column.
	//
	// During a dataset update, if the column ID
	//             of a calculated column matches that of an existing calculated column, Amazon QuickSight
	//             preserves the existing calculated column.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#column_id QuicksightDataSet#column_id}
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// <p>Column name.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// <p>An expression that defines the calculated column.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#expression QuicksightDataSet#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

