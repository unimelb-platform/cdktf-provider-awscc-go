package lexbot


type LexBotBotLocalesSlotTypesExternalSourceSettingGrammarSlotTypeSetting struct {
	// Describes the Amazon S3 bucket name and location for the grammar that is the source for the slot type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#source LexBot#source}
	Source *LexBotBotLocalesSlotTypesExternalSourceSettingGrammarSlotTypeSettingSource `field:"optional" json:"source" yaml:"source"`
}

