package quicksighttemplate


type QuicksightTemplateSourceEntitySourceAnalysis struct {
	// <p>The Amazon Resource Name (ARN) of the resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#arn QuicksightTemplate#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// <p>A structure containing information about the dataset references used as placeholders             in the template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#data_set_references QuicksightTemplate#data_set_references}
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

