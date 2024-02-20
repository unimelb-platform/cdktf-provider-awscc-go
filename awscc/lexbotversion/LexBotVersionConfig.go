package lexbotversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LexBotVersionConfig struct {
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
	// Unique ID of resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_version#bot_id LexBotVersion#bot_id}
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Specifies the locales that Amazon Lex adds to this version.
	//
	// You can choose the Draft version or any other previously published version for each locale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_version#bot_version_locale_specification LexBotVersion#bot_version_locale_specification}
	BotVersionLocaleSpecification interface{} `field:"required" json:"botVersionLocaleSpecification" yaml:"botVersionLocaleSpecification"`
	// A description of the version. Use the description to help identify the version in lists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot_version#description LexBotVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

