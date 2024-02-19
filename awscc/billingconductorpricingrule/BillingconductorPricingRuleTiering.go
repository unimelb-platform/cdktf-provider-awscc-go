package billingconductorpricingrule


type BillingconductorPricingRuleTiering struct {
	// The possible customizable free tier configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#free_tier BillingconductorPricingRule#free_tier}
	FreeTier *BillingconductorPricingRuleTieringFreeTier `field:"optional" json:"freeTier" yaml:"freeTier"`
}

