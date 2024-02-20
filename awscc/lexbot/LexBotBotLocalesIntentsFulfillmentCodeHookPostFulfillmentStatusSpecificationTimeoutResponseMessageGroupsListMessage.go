package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseMessageGroupsListMessage struct {
	// A message in a custom format defined by the client application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#custom_payload LexBot#custom_payload}
	CustomPayload *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseMessageGroupsListMessageCustomPayload `field:"optional" json:"customPayload" yaml:"customPayload"`
	// A message that defines a response card that the client application can show to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#image_response_card LexBot#image_response_card}
	ImageResponseCard *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseMessageGroupsListMessageImageResponseCard `field:"optional" json:"imageResponseCard" yaml:"imageResponseCard"`
	// A message in plain text format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#plain_text_message LexBot#plain_text_message}
	PlainTextMessage *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseMessageGroupsListMessagePlainTextMessage `field:"optional" json:"plainTextMessage" yaml:"plainTextMessage"`
	// A message in Speech Synthesis Markup Language (SSML).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#ssml_message LexBot#ssml_message}
	SsmlMessage *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseMessageGroupsListMessageSsmlMessage `field:"optional" json:"ssmlMessage" yaml:"ssmlMessage"`
}

