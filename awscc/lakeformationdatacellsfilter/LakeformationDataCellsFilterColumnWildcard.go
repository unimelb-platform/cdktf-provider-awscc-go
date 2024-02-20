package lakeformationdatacellsfilter


type LakeformationDataCellsFilterColumnWildcard struct {
	// A list of column names to be excluded from the Data Cells Filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lakeformation_data_cells_filter#excluded_column_names LakeformationDataCellsFilter#excluded_column_names}
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

