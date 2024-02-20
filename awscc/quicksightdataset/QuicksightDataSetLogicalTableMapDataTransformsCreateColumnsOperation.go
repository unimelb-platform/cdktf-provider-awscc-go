package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation struct {
	// <p>Calculated columns to create.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
}

