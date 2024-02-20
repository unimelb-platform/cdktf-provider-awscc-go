package lexbotalias


type LexBotAliasBotAliasLocaleSettingsBotAliasLocaleSetting struct {
	// Whether the Lambda code hook is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#enabled LexBotAlias#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Contains information about code hooks that Amazon Lex calls during a conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#code_hook_specification LexBotAlias#code_hook_specification}
	CodeHookSpecification *LexBotAliasBotAliasLocaleSettingsBotAliasLocaleSettingCodeHookSpecification `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
}

