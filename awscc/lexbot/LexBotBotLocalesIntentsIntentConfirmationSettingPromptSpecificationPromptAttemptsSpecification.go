package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecification struct {
	// Specifies the allowed input types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#allowed_input_types LexBot#allowed_input_types}
	AllowedInputTypes *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAllowedInputTypes `field:"optional" json:"allowedInputTypes" yaml:"allowedInputTypes"`
	// Indicates whether the user can interrupt a speech prompt attempt from the bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#allow_interrupt LexBot#allow_interrupt}
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Specifies the audio and DTMF input specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#audio_and_dtmf_input_specification LexBot#audio_and_dtmf_input_specification}
	AudioAndDtmfInputSpecification *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecification `field:"optional" json:"audioAndDtmfInputSpecification" yaml:"audioAndDtmfInputSpecification"`
	// Specifies the text input specifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#text_input_specification LexBot#text_input_specification}
	TextInputSpecification *LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationTextInputSpecification `field:"optional" json:"textInputSpecification" yaml:"textInputSpecification"`
}

