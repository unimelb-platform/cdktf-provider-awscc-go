package quicksightdashboard


type QuicksightDashboardParametersStringParameters struct {
	// <p>A display name for a string parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#name QuicksightDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The values of a string parameter.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#values QuicksightDashboard#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

