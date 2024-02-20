package quicksightanalysis


type QuicksightAnalysisTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#key QuicksightAnalysis#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#value QuicksightAnalysis#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

