package apsrulegroupsnamespace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApsRuleGroupsNamespaceConfig struct {
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
	// The RuleGroupsNamespace data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_rule_groups_namespace#data ApsRuleGroupsNamespace#data}
	Data *string `field:"required" json:"data" yaml:"data"`
	// The RuleGroupsNamespace name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_rule_groups_namespace#name ApsRuleGroupsNamespace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required to identify a specific APS Workspace associated with this RuleGroupsNamespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_rule_groups_namespace#workspace ApsRuleGroupsNamespace#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_rule_groups_namespace#tags ApsRuleGroupsNamespace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

