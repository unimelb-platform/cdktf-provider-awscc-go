package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationUpdateResponseMessageGroupsVariationsImageResponseCardButtons struct {
	// The text that appears on the button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lex_bot#text LexBot#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The value returned to Amazon Lex when the user chooses this button.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lex_bot#value LexBot#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

