package lexbot


type LexBotTestBotAliasSettingsConversationLogSettings struct {
	// List of audio log settings that pertain to the conversation log settings for the bot's TestBotAlias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#audio_log_settings LexBot#audio_log_settings}
	AudioLogSettings interface{} `field:"optional" json:"audioLogSettings" yaml:"audioLogSettings"`
	// List of text log settings that pertain to the conversation log settings for the bot's TestBotAlias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#text_log_settings LexBot#text_log_settings}
	TextLogSettings interface{} `field:"optional" json:"textLogSettings" yaml:"textLogSettings"`
}

