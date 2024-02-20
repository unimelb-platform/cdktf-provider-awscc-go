package billingconductorcustomlineitem


type BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#attribute BillingconductorCustomLineItem#attribute}.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#match_option BillingconductorCustomLineItem#match_option}.
	MatchOption *string `field:"required" json:"matchOption" yaml:"matchOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#values BillingconductorCustomLineItem#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

