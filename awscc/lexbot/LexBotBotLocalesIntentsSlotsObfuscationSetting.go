package lexbot


type LexBotBotLocalesIntentsSlotsObfuscationSetting struct {
	// Value that determines whether Amazon Lex obscures slot values in conversation logs. The default is to obscure the values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#obfuscation_setting_type LexBot#obfuscation_setting_type}
	ObfuscationSettingType *string `field:"required" json:"obfuscationSettingType" yaml:"obfuscationSettingType"`
}

