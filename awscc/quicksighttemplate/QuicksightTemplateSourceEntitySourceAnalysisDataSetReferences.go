package quicksighttemplate


type QuicksightTemplateSourceEntitySourceAnalysisDataSetReferences struct {
	// <p>Dataset Amazon Resource Name (ARN).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#data_set_arn QuicksightTemplate#data_set_arn}
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// <p>Dataset placeholder.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#data_set_placeholder QuicksightTemplate#data_set_placeholder}
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

