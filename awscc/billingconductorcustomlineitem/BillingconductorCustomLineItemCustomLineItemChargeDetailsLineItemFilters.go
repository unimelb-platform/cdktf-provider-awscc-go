package billingconductorcustomlineitem


type BillingconductorCustomLineItemCustomLineItemChargeDetailsLineItemFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/billingconductor_custom_line_item#attribute BillingconductorCustomLineItem#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/billingconductor_custom_line_item#match_option BillingconductorCustomLineItem#match_option}.
	MatchOption *string `field:"optional" json:"matchOption" yaml:"matchOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/billingconductor_custom_line_item#values BillingconductorCustomLineItem#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

