package wisdomaiguardrail


type WisdomAiGuardrailContextualGroundingPolicyConfig struct {
	// List of contextual grounding filter configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#filters_config WisdomAiGuardrail#filters_config}
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
}

