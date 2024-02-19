package quicksightanalysis


type QuicksightAnalysisParameters struct {
	// <p>Date-time parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#date_time_parameters QuicksightAnalysis#date_time_parameters}
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// <p>Decimal parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#decimal_parameters QuicksightAnalysis#decimal_parameters}
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// <p>Integer parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#integer_parameters QuicksightAnalysis#integer_parameters}
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// <p>String parameters.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_analysis#string_parameters QuicksightAnalysis#string_parameters}
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

