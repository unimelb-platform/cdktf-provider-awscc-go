package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSetting struct {
	// A list of message groups that Amazon Lex uses to respond the user input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#declination_response LexBot#declination_response}
	DeclinationResponse *LexBotBotLocalesIntentsIntentConfirmationSettingDeclinationResponse `field:"required" json:"declinationResponse" yaml:"declinationResponse"`
	// Prompts the user to confirm the intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#prompt_specification LexBot#prompt_specification}
	PromptSpecification *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecification `field:"required" json:"promptSpecification" yaml:"promptSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#is_active LexBot#is_active}.
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
}

