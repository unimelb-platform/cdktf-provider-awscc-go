package cecostcategory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CeCostCategoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_cost_category#name CeCostCategory#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// JSON array format of Expression in Billing and Cost Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_cost_category#rules CeCostCategory#rules}
	Rules *string `field:"required" json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_cost_category#rule_version CeCostCategory#rule_version}.
	RuleVersion *string `field:"required" json:"ruleVersion" yaml:"ruleVersion"`
	// The default value for the cost category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_cost_category#default_value CeCostCategory#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Json array format of CostCategorySplitChargeRule in Billing and Cost Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_cost_category#split_charge_rules CeCostCategory#split_charge_rules}
	SplitChargeRules *string `field:"optional" json:"splitChargeRules" yaml:"splitChargeRules"`
}

