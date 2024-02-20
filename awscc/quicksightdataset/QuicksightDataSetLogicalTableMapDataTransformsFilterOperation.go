package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsFilterOperation struct {
	// <p>An expression that must evaluate to a Boolean value.
	//
	// Rows for which the expression
	//             evaluates to true are kept in the dataset.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#condition_expression QuicksightDataSet#condition_expression}
	ConditionExpression *string `field:"required" json:"conditionExpression" yaml:"conditionExpression"`
}

