package billingconductorcustomlineitem


type BillingconductorCustomLineItemCustomLineItemChargeDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#type BillingconductorCustomLineItem#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#flat BillingconductorCustomLineItem#flat}.
	Flat *BillingconductorCustomLineItemCustomLineItemChargeDetailsFlat `field:"optional" json:"flat" yaml:"flat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#line_item_filters BillingconductorCustomLineItem#line_item_filters}.
	LineItemFilters interface{} `field:"optional" json:"lineItemFilters" yaml:"lineItemFilters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#percentage BillingconductorCustomLineItem#percentage}.
	Percentage *BillingconductorCustomLineItemCustomLineItemChargeDetailsPercentage `field:"optional" json:"percentage" yaml:"percentage"`
}

