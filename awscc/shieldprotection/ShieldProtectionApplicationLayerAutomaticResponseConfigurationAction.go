package shieldprotection


type ShieldProtectionApplicationLayerAutomaticResponseConfigurationAction struct {
	// Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.
	//
	// You must specify exactly one action, either `Block` or `Count`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#block ShieldProtection#block}
	Block *string `field:"optional" json:"block" yaml:"block"`
	// Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.
	//
	// You must specify exactly one action, either `Block` or `Count`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#count ShieldProtection#count}
	Count *string `field:"optional" json:"count" yaml:"count"`
}

