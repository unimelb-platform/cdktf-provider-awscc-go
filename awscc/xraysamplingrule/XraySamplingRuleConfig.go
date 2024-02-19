package xraysamplingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type XraySamplingRuleConfig struct {
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
	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#rule_name XraySamplingRule#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#sampling_rule XraySamplingRule#sampling_rule}.
	SamplingRule *XraySamplingRuleSamplingRule `field:"optional" json:"samplingRule" yaml:"samplingRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#sampling_rule_record XraySamplingRule#sampling_rule_record}.
	SamplingRuleRecord *XraySamplingRuleSamplingRuleRecord `field:"optional" json:"samplingRuleRecord" yaml:"samplingRuleRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#sampling_rule_update XraySamplingRule#sampling_rule_update}.
	SamplingRuleUpdate *XraySamplingRuleSamplingRuleUpdate `field:"optional" json:"samplingRuleUpdate" yaml:"samplingRuleUpdate"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_sampling_rule#tags XraySamplingRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

