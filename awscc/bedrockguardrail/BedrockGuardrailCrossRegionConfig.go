package bedrockguardrail


type BedrockGuardrailCrossRegionConfig struct {
	// The Amazon Resource Name (ARN) of the guardrail profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#guardrail_profile_arn BedrockGuardrail#guardrail_profile_arn}
	GuardrailProfileArn *string `field:"optional" json:"guardrailProfileArn" yaml:"guardrailProfileArn"`
}

