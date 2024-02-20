package quicksightdataset


type QuicksightDataSetPhysicalTableMapS3SourceInputColumns struct {
	// <p>The name of this column in the underlying data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#type QuicksightDataSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

