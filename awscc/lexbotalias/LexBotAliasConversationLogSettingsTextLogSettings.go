package lexbotalias


type LexBotAliasConversationLogSettingsTextLogSettings struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#destination LexBotAlias#destination}
	Destination *LexBotAliasConversationLogSettingsTextLogSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#enabled LexBotAlias#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

