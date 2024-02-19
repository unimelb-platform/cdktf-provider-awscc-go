package lexbot


type LexBotBotLocalesIntentsSampleUtterances struct {
	// The sample utterance that Amazon Lex uses to build its machine-learning model to recognize intents/slots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#utterance LexBot#utterance}
	Utterance *string `field:"required" json:"utterance" yaml:"utterance"`
}

