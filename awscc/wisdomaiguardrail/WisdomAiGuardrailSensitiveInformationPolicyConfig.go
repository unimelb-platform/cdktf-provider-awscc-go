package wisdomaiguardrail


type WisdomAiGuardrailSensitiveInformationPolicyConfig struct {
	// List of entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#pii_entities_config WisdomAiGuardrail#pii_entities_config}
	PiiEntitiesConfig interface{} `field:"optional" json:"piiEntitiesConfig" yaml:"piiEntitiesConfig"`
	// List of regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_guardrail#regexes_config WisdomAiGuardrail#regexes_config}
	RegexesConfig interface{} `field:"optional" json:"regexesConfig" yaml:"regexesConfig"`
}

