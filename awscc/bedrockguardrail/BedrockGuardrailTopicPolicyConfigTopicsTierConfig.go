package bedrockguardrail


type BedrockGuardrailTopicPolicyConfigTopicsTierConfig struct {
	// Tier name for tier configuration in topic policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#tier_name BedrockGuardrail#tier_name}
	TierName *string `field:"optional" json:"tierName" yaml:"tierName"`
}

