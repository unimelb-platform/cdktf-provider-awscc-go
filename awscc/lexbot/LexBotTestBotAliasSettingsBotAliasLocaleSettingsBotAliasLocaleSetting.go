package lexbot


type LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSetting struct {
	// Whether the Lambda code hook is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#enabled LexBot#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Contains information about code hooks that Amazon Lex calls during a conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#code_hook_specification LexBot#code_hook_specification}
	CodeHookSpecification *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecification `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
}

