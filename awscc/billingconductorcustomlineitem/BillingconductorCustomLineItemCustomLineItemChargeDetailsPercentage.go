package billingconductorcustomlineitem


type BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#percentage_value BillingconductorCustomLineItem#percentage_value}.
	PercentageValue *float64 `field:"required" json:"percentageValue" yaml:"percentageValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#child_associated_resources BillingconductorCustomLineItem#child_associated_resources}.
	ChildAssociatedResources *[]*string `field:"optional" json:"childAssociatedResources" yaml:"childAssociatedResources"`
}

