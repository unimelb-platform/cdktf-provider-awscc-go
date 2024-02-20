package lexbot


type LexBotBotLocalesIntentsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecification struct {
	// The DTMF character that clears the accumulated DTMF digits and immediately ends the input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#deletion_character LexBot#deletion_character}
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// The DTMF character that immediately ends input.
	//
	// If the user does not press this character, the input ends after the end timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#end_character LexBot#end_character}
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// How long the bot should wait after the last DTMF character input before assuming that the input has concluded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#end_timeout_ms LexBot#end_timeout_ms}
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// The maximum number of DTMF digits allowed in an utterance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#max_length LexBot#max_length}
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

