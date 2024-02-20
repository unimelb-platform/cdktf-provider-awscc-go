package databrewruleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabrewRulesetConfig struct {
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
	// Name of the Ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#name DatabrewRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of the data quality rules in the ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#rules DatabrewRuleset#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// Arn of the target resource (dataset) to apply the ruleset to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#target_arn DatabrewRuleset#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Description of the Ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#description DatabrewRuleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_ruleset#tags DatabrewRuleset#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

