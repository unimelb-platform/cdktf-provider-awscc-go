package bedrockguardrail


type BedrockGuardrailTags struct {
	// Tag Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#key BedrockGuardrail#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#value BedrockGuardrail#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

