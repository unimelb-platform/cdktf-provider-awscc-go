package quicksightdashboard


type QuicksightDashboardSourceEntitySourceTemplate struct {
	// <p>The Amazon Resource Name (ARN) of the resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#arn QuicksightDashboard#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// <p>Dataset references.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#data_set_references QuicksightDashboard#data_set_references}
	DataSetReferences interface{} `field:"required" json:"dataSetReferences" yaml:"dataSetReferences"`
}

