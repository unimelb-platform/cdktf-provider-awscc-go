package quicksightdatasource


type QuicksightDataSourceTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#key QuicksightDataSource#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#value QuicksightDataSource#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

