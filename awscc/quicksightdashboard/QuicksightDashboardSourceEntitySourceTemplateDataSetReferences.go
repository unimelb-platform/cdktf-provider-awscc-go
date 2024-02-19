package quicksightdashboard


type QuicksightDashboardSourceEntitySourceTemplateDataSetReferences struct {
	// <p>Dataset Amazon Resource Name (ARN).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#data_set_arn QuicksightDashboard#data_set_arn}
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// <p>Dataset placeholder.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#data_set_placeholder QuicksightDashboard#data_set_placeholder}
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

