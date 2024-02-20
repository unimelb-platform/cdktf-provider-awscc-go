package route53resolverfirewallrulegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverFirewallRuleGroupConfig struct {
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
	// FirewallRules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_firewall_rule_group#firewall_rules Route53ResolverFirewallRuleGroup#firewall_rules}
	FirewallRules interface{} `field:"optional" json:"firewallRules" yaml:"firewallRules"`
	// FirewallRuleGroupName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_firewall_rule_group#name Route53ResolverFirewallRuleGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53resolver_firewall_rule_group#tags Route53ResolverFirewallRuleGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

