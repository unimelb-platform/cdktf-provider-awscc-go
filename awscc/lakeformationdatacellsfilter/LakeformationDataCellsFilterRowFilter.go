package lakeformationdatacellsfilter


type LakeformationDataCellsFilterRowFilter struct {
	// An empty object representing a row wildcard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_data_cells_filter#all_rows_wildcard LakeformationDataCellsFilter#all_rows_wildcard}
	AllRowsWildcard *string `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// A PartiQL predicate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_data_cells_filter#filter_expression LakeformationDataCellsFilter#filter_expression}
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

