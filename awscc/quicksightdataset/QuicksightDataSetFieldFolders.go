package quicksightdataset


type QuicksightDataSetFieldFolders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#description QuicksightDataSet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

