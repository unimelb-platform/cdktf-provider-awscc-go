package lexbot


type LexBotBotLocalesVoiceSettings struct {
	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#voice_id LexBot#voice_id}
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
	// Indicates the type of Amazon Polly voice that Amazon Lex should use for voice interaction with the user.
	//
	// For more information, see the engine parameter of the SynthesizeSpeech operation in the Amazon Polly developer guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#engine LexBot#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
}

