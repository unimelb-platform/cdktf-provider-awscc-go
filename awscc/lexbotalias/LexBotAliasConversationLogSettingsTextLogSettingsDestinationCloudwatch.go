package lexbotalias


type LexBotAliasConversationLogSettingsTextLogSettingsDestinationCloudwatch struct {
	// A string used to identify the groupArn for the Cloudwatch Log Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#cloudwatch_log_group_arn LexBotAlias#cloudwatch_log_group_arn}
	CloudwatchLogGroupArn *string `field:"required" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// A string containing the value for the Log Prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#log_prefix LexBotAlias#log_prefix}
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
}

