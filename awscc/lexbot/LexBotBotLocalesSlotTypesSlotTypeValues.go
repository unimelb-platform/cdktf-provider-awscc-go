package lexbot


type LexBotBotLocalesSlotTypesSlotTypeValues struct {
	// Defines one of the values for a slot type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#sample_value LexBot#sample_value}
	SampleValue *LexBotBotLocalesSlotTypesSlotTypeValuesSampleValue `field:"required" json:"sampleValue" yaml:"sampleValue"`
	// Additional values related to the slot type entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#synonyms LexBot#synonyms}
	Synonyms interface{} `field:"optional" json:"synonyms" yaml:"synonyms"`
}

