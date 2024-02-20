package lexbot


type LexBotBotLocalesIntentsSlotsValueElicitationSettingPromptSpecification struct {
	// The maximum number of times the bot tries to elicit a resonse from the user using this prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#max_retries LexBot#max_retries}
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// One to 5 message groups that contain update messages.
	//
	// Amazon Lex chooses one of the messages to play to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#message_groups_list LexBot#message_groups_list}
	MessageGroupsList interface{} `field:"required" json:"messageGroupsList" yaml:"messageGroupsList"`
	// Indicates whether the user can interrupt a speech prompt from the bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#allow_interrupt LexBot#allow_interrupt}
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Indicates how a message is selected from a message group among retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#message_selection_strategy LexBot#message_selection_strategy}
	MessageSelectionStrategy *string `field:"optional" json:"messageSelectionStrategy" yaml:"messageSelectionStrategy"`
	// Specifies the advanced settings on each attempt of the prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#prompt_attempts_specification LexBot#prompt_attempts_specification}
	PromptAttemptsSpecification interface{} `field:"optional" json:"promptAttemptsSpecification" yaml:"promptAttemptsSpecification"`
}

