package lexbot


type LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSetting struct {
	// Contains information about code hooks that Amazon Lex calls during a conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lex_bot#code_hook_specification LexBot#code_hook_specification}
	CodeHookSpecification *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecification `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
	// Whether the Lambda code hook is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lex_bot#enabled LexBot#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

