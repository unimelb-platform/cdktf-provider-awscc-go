package wisdomaiguardrail


type WisdomAiGuardrailSensitiveInformationPolicyConfigPiiEntitiesConfig struct {
	// Options for sensitive information action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#action WisdomAiGuardrail#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The currently supported PII entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#type WisdomAiGuardrail#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

