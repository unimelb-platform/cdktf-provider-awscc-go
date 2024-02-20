package lexbot


type LexBotTestBotAliasSettingsSentimentAnalysisSettings struct {
	// Enable to call Amazon Comprehend for Sentiment natively within Lex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#detect_sentiment LexBot#detect_sentiment}
	DetectSentiment interface{} `field:"required" json:"detectSentiment" yaml:"detectSentiment"`
}

