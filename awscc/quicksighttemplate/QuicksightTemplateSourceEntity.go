package quicksighttemplate


type QuicksightTemplateSourceEntity struct {
	// <p>The source analysis of the template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#source_analysis QuicksightTemplate#source_analysis}
	SourceAnalysis *QuicksightTemplateSourceEntitySourceAnalysis `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// <p>The source template of the template.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#source_template QuicksightTemplate#source_template}
	SourceTemplate *QuicksightTemplateSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

