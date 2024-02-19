package lexbotalias


type LexBotAliasConversationLogSettingsAudioLogSettingsDestination struct {
	// Specifies an Amazon S3 bucket for logging audio conversations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#s3_bucket LexBotAlias#s3_bucket}
	S3Bucket *LexBotAliasConversationLogSettingsAudioLogSettingsDestinationS3Bucket `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

