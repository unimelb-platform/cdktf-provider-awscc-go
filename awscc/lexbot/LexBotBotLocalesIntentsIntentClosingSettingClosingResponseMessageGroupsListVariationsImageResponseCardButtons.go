package lexbot


type LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtons struct {
	// The text that appears on the button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#text LexBot#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#value LexBot#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

