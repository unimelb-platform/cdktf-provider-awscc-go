package quicksightanalysis


type QuicksightAnalysisSourceEntity struct {
	// <p>The source template of an analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#source_template QuicksightAnalysis#source_template}
	SourceTemplate *QuicksightAnalysisSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

