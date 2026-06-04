package wisdomaiguardrail


type WisdomAiGuardrailContentPolicyConfigFiltersConfig struct {
	// Strength for filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#input_strength WisdomAiGuardrail#input_strength}
	InputStrength *string `field:"optional" json:"inputStrength" yaml:"inputStrength"`
	// Strength for filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#output_strength WisdomAiGuardrail#output_strength}
	OutputStrength *string `field:"optional" json:"outputStrength" yaml:"outputStrength"`
	// Type of text to text filter in content policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#type WisdomAiGuardrail#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

