package quicksightanalysis


type QuicksightAnalysisSourceEntitySourceTemplateDataSetReferences struct {
	// <p>Dataset Amazon Resource Name (ARN).</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#data_set_arn QuicksightAnalysis#data_set_arn}
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// <p>Dataset placeholder.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#data_set_placeholder QuicksightAnalysis#data_set_placeholder}
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

