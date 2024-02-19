package lexbotalias


type LexBotAliasSentimentAnalysisSettings struct {
	// Enable to call Amazon Comprehend for Sentiment natively within Lex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#detect_sentiment LexBotAlias#detect_sentiment}
	DetectSentiment interface{} `field:"required" json:"detectSentiment" yaml:"detectSentiment"`
}

