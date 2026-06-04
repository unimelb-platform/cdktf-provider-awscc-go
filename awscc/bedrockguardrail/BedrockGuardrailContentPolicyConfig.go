package bedrockguardrail


type BedrockGuardrailContentPolicyConfig struct {
	// Guardrail tier config for content policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#content_filters_tier_config BedrockGuardrail#content_filters_tier_config}
	ContentFiltersTierConfig *BedrockGuardrailContentPolicyConfigContentFiltersTierConfig `field:"optional" json:"contentFiltersTierConfig" yaml:"contentFiltersTierConfig"`
	// List of content filter configs in content policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#filters_config BedrockGuardrail#filters_config}
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
}

