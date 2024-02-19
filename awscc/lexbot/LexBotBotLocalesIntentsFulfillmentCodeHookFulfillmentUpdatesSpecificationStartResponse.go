package lexbot


type LexBotBotLocalesIntentsFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponse struct {
	// The delay between when the Lambda fulfillment function starts running and the start message is played.
	//
	// If the Lambda function returns before the delay is over, the start message isn't played.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#delay_in_seconds LexBot#delay_in_seconds}
	DelayInSeconds *float64 `field:"required" json:"delayInSeconds" yaml:"delayInSeconds"`
	// One to 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#message_groups LexBot#message_groups}
	MessageGroups interface{} `field:"required" json:"messageGroups" yaml:"messageGroups"`
	// Determines whether the user can interrupt the start message while it is playing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#allow_interrupt LexBot#allow_interrupt}
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
}

