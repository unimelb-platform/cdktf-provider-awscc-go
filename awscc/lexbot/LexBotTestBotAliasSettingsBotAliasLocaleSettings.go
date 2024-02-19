package lexbot


type LexBotTestBotAliasSettingsBotAliasLocaleSettings struct {
	// You can use this parameter to specify a specific Lambda function to run different functions in different locales.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#bot_alias_locale_setting LexBot#bot_alias_locale_setting}
	BotAliasLocaleSetting *LexBotTestBotAliasSettingsBotAliasLocaleSettingsBotAliasLocaleSetting `field:"required" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// A string used to identify the locale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#locale_id LexBot#locale_id}
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

