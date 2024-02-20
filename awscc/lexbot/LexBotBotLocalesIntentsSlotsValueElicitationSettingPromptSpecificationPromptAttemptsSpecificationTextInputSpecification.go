package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecification struct {
	// Time for which a bot waits before re-prompting a customer for text input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#start_timeout_ms LexBot#start_timeout_ms}
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
}

