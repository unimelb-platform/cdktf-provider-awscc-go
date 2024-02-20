package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsTextLogSettings struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#destination LexBot#destination}
	Destination *LexBotTestBotAliasSettingsConversationLogSettingsTextLogSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#enabled LexBot#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

