package lexbot


type LexBotBotLocalesSlotTypesValueSelectionSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#resolution_strategy LexBot#resolution_strategy}.
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// Provides settings that enable advanced recognition settings for slot values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#advanced_recognition_setting LexBot#advanced_recognition_setting}
	AdvancedRecognitionSetting *LexBotBotLocalesSlotTypesValueSelectionSettingAdvancedRecognitionSetting `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// A regular expression used to validate the value of a slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#regex_filter LexBot#regex_filter}
	RegexFilter *LexBotBotLocalesSlotTypesValueSelectionSettingRegexFilter `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}

