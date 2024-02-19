package quicksightdatasource


type QuicksightDataSourceErrorInfo struct {
	// <p>Error message.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#message QuicksightDataSource#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#type QuicksightDataSource#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

