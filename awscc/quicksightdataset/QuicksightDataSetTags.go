package quicksightdataset


type QuicksightDataSetTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#key QuicksightDataSet#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#value QuicksightDataSet#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

