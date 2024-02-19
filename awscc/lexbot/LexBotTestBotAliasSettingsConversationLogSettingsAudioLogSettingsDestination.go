package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettingsDestination struct {
	// Specifies an Amazon S3 bucket for logging audio conversations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#s3_bucket LexBot#s3_bucket}
	S3Bucket *LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettingsDestinationS3Bucket `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

