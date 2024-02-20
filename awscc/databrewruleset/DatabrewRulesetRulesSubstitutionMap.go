package databrewruleset


type DatabrewRulesetRulesSubstitutionMap struct {
	// Value or column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#value DatabrewRuleset#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#value_reference DatabrewRuleset#value_reference}
	ValueReference *string `field:"required" json:"valueReference" yaml:"valueReference"`
}

