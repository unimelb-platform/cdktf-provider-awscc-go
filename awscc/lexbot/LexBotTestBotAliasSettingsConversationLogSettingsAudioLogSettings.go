package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettings struct {
	// The location of audio log files collected when conversation logging is enabled for a bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#destination LexBot#destination}
	Destination *LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#enabled LexBot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

