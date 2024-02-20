package route53recoverycontrolsafetyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoverycontrolSafetyRuleConfig struct {
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
	// An assertion rule enforces that, when a routing control state is changed, that the criteria set by the rule configuration is met.
	//
	// Otherwise, the change to the routing control is not accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule#assertion_rule Route53RecoverycontrolSafetyRule#assertion_rule}
	AssertionRule *Route53RecoverycontrolSafetyRuleAssertionRule `field:"optional" json:"assertionRule" yaml:"assertionRule"`
	// The Amazon Resource Name (ARN) of the control panel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule#control_panel_arn Route53RecoverycontrolSafetyRule#control_panel_arn}
	ControlPanelArn *string `field:"optional" json:"controlPanelArn" yaml:"controlPanelArn"`
	// A gating rule verifies that a set of gating controls evaluates as true, based on a rule configuration that you specify.
	//
	// If the gating rule evaluates to true, Amazon Route 53 Application Recovery Controller allows a set of routing control state changes to run and complete against the set of target controls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule#gating_rule Route53RecoverycontrolSafetyRule#gating_rule}
	GatingRule *Route53RecoverycontrolSafetyRuleGatingRule `field:"optional" json:"gatingRule" yaml:"gatingRule"`
	// The name for the safety rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule#name Route53RecoverycontrolSafetyRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The rule configuration for an assertion rule or gating rule.
	//
	// This is the criteria that you set for specific assertion controls (routing controls) or gating controls. This configuration specifies how many controls must be enabled after a transaction completes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule#rule_config Route53RecoverycontrolSafetyRule#rule_config}
	RuleConfig *Route53RecoverycontrolSafetyRuleRuleConfig `field:"optional" json:"ruleConfig" yaml:"ruleConfig"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_safety_rule#tags Route53RecoverycontrolSafetyRule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

