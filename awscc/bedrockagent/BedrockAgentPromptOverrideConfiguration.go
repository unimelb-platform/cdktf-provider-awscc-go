package bedrockagent


type BedrockAgentPromptOverrideConfiguration struct {
	// ARN of a Lambda.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent#override_lambda BedrockAgent#override_lambda}
	OverrideLambda *string `field:"optional" json:"overrideLambda" yaml:"overrideLambda"`
	// List of BasePromptConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent#prompt_configurations BedrockAgent#prompt_configurations}
	PromptConfigurations interface{} `field:"optional" json:"promptConfigurations" yaml:"promptConfigurations"`
}

