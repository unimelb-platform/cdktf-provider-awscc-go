package quicksightanalysis


type QuicksightAnalysisSourceEntitySourceTemplate struct {
	// <p>The Amazon Resource Name (ARN) of the source template of an analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#arn QuicksightAnalysis#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// <p>The dataset references of the source template of an analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#data_set_references QuicksightAnalysis#data_set_references}
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

