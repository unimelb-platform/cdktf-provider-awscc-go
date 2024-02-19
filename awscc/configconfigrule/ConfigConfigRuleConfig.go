package configconfigrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConfigRuleConfig struct {
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
	// Source of events for the AWS Config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#source ConfigConfigRule#source}
	Source *ConfigConfigRuleSource `field:"required" json:"source" yaml:"source"`
	// Compliance details of the Config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#compliance ConfigConfigRule#compliance}
	Compliance *ConfigConfigRuleCompliance `field:"optional" json:"compliance" yaml:"compliance"`
	// Name for the AWS Config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#config_rule_name ConfigConfigRule#config_rule_name}
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// Description provided for the AWS Config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#description ConfigConfigRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of EvaluationModeConfiguration objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#evaluation_modes ConfigConfigRule#evaluation_modes}
	EvaluationModes interface{} `field:"optional" json:"evaluationModes" yaml:"evaluationModes"`
	// JSON string passed the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#input_parameters ConfigConfigRule#input_parameters}
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// Maximum frequency at which the rule has to be evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Scope to constrain which resources can trigger the AWS Config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#scope ConfigConfigRule#scope}
	Scope *ConfigConfigRuleScope `field:"optional" json:"scope" yaml:"scope"`
}

