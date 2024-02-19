package quicksightanalysis


type QuicksightAnalysisParametersDateTimeParameters struct {
	// <p>A display name for the date-time parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The values for the date-time parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#values QuicksightAnalysis#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

