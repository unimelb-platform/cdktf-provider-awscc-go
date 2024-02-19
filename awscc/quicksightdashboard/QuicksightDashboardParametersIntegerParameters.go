package quicksightdashboard


type QuicksightDashboardParametersIntegerParameters struct {
	// <p>The name of the integer parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The values for the integer parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

