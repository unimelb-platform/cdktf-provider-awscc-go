package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#enabled LexBot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Provides information for updating the user on the progress of fulfilling an intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#fulfillment_updates_specification LexBot#fulfillment_updates_specification}
	FulfillmentUpdatesSpecification *LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecification `field:"optional" json:"fulfillmentUpdatesSpecification" yaml:"fulfillmentUpdatesSpecification"`
	// Provides information for updating the user on the progress of fulfilling an intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#post_fulfillment_status_specification LexBot#post_fulfillment_status_specification}
	PostFulfillmentStatusSpecification *LexBotBotLocalesIntentsFulfillmentCodeHookPostFulfillmentStatusSpecification `field:"optional" json:"postFulfillmentStatusSpecification" yaml:"postFulfillmentStatusSpecification"`
}

