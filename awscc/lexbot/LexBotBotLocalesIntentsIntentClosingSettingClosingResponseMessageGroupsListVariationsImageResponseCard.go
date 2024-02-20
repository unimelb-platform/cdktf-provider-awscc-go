package lexbot


type LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCard struct {
	// The title to display on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#title LexBot#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of buttons that should be displayed on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#buttons LexBot#buttons}
	Buttons interface{} `field:"optional" json:"buttons" yaml:"buttons"`
	// The URL of an image to display on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#image_url LexBot#image_url}
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The subtitle to display on the response card.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#subtitle LexBot#subtitle}
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

