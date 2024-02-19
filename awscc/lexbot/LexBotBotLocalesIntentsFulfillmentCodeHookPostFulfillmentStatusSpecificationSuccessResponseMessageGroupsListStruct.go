package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListStruct struct {
	// The primary message that Amazon Lex should send to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#message LexBot#message}
	Message *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseMessageGroupsListMessage `field:"required" json:"message" yaml:"message"`
	// Message variations to send to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#variations LexBot#variations}
	Variations interface{} `field:"optional" json:"variations" yaml:"variations"`
}

