package lexbotalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LexBotAliasConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A unique identifier for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#bot_alias_name LexBotAlias#bot_alias_name}
	BotAliasName *string `field:"required" json:"botAliasName" yaml:"botAliasName"`
	// Unique ID of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#bot_id LexBotAlias#bot_id}
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// A list of bot alias locale settings to add to the bot alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#bot_alias_locale_settings LexBotAlias#bot_alias_locale_settings}
	BotAliasLocaleSettings interface{} `field:"optional" json:"botAliasLocaleSettings" yaml:"botAliasLocaleSettings"`
	// A list of tags to add to the bot alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#bot_alias_tags LexBotAlias#bot_alias_tags}
	BotAliasTags interface{} `field:"optional" json:"botAliasTags" yaml:"botAliasTags"`
	// The version of a bot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#bot_version LexBotAlias#bot_version}
	BotVersion *string `field:"optional" json:"botVersion" yaml:"botVersion"`
	// Contains information about code hooks that Amazon Lex calls during a conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#conversation_log_settings LexBotAlias#conversation_log_settings}
	ConversationLogSettings *LexBotAliasConversationLogSettings `field:"optional" json:"conversationLogSettings" yaml:"conversationLogSettings"`
	// A description of the bot alias. Use the description to help identify the bot alias in lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#description LexBotAlias#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_alias#sentiment_analysis_settings LexBotAlias#sentiment_analysis_settings}
	SentimentAnalysisSettings *LexBotAliasSentimentAnalysisSettings `field:"optional" json:"sentimentAnalysisSettings" yaml:"sentimentAnalysisSettings"`
}

