package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsProjectOperation struct {
	// <p>Projected columns.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#projected_columns QuicksightDataSet#projected_columns}
	ProjectedColumns *[]*string `field:"required" json:"projectedColumns" yaml:"projectedColumns"`
}

