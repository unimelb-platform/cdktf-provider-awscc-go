package billingconductorbillinggroup


type BillingconductorBillingGroupComputationPreference struct {
	// ARN of the attached pricing plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_billing_group#pricing_plan_arn BillingconductorBillingGroup#pricing_plan_arn}
	PricingPlanArn *string `field:"required" json:"pricingPlanArn" yaml:"pricingPlanArn"`
}

