package billingconductorbillinggroup


type BillingconductorBillingGroupAccountGrouping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_billing_group#linked_account_ids BillingconductorBillingGroup#linked_account_ids}.
	LinkedAccountIds *[]*string `field:"required" json:"linkedAccountIds" yaml:"linkedAccountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_billing_group#auto_associate BillingconductorBillingGroup#auto_associate}.
	AutoAssociate interface{} `field:"optional" json:"autoAssociate" yaml:"autoAssociate"`
}

