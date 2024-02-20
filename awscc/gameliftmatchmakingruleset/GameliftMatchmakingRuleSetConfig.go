package gameliftmatchmakingruleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GameliftMatchmakingRuleSetConfig struct {
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
	// A unique identifier for the matchmaking rule set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_rule_set#name GameliftMatchmakingRuleSet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A collection of matchmaking rules, formatted as a JSON string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_rule_set#rule_set_body GameliftMatchmakingRuleSet#rule_set_body}
	RuleSetBody *string `field:"required" json:"ruleSetBody" yaml:"ruleSetBody"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_rule_set#tags GameliftMatchmakingRuleSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

