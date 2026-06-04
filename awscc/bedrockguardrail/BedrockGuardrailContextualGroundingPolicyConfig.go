package bedrockguardrail


type BedrockGuardrailContextualGroundingPolicyConfig struct {
	// List of contextual grounding filter configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#filters_config BedrockGuardrail#filters_config}
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
}

