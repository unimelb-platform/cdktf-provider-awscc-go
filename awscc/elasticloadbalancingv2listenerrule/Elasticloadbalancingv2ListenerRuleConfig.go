package elasticloadbalancingv2listenerrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Elasticloadbalancingv2ListenerRuleConfig struct {
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
	// The actions.
	//
	// The rule must include exactly one of the following types of actions: ``forward``, ``fixed-response``, or ``redirect``, and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener_rule#actions Elasticloadbalancingv2ListenerRule#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions.
	//
	// The rule can optionally include up to one of each of the following conditions: ``http-request-method``, ``host-header``, ``path-pattern``, and ``source-ip``. A rule can also optionally include one or more of each of the following conditions: ``http-header`` and ``query-string``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener_rule#conditions Elasticloadbalancingv2ListenerRule#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The rule priority.
	//
	// A listener can't have multiple rules with the same priority.
	//  If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener_rule#priority Elasticloadbalancingv2ListenerRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The Amazon Resource Name (ARN) of the listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/elasticloadbalancingv2_listener_rule#listener_arn Elasticloadbalancingv2ListenerRule#listener_arn}
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
}

