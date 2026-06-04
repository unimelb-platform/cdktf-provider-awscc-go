package wisdomaiguardrail


type WisdomAiGuardrailTopicPolicyConfig struct {
	// List of topic configs in topic policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#topics_config WisdomAiGuardrail#topics_config}
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
}

