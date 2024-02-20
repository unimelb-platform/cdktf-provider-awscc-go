package quicksightanalysis


type QuicksightAnalysisErrors struct {
	// <p>The message associated with the analysis error.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#message QuicksightAnalysis#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#type QuicksightAnalysis#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

