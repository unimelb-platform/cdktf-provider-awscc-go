package billingconductorpricingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingconductorPricingRuleConfig struct {
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
	// Pricing rule name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#name BillingconductorPricingRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A term used to categorize the granularity of a Pricing Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#scope BillingconductorPricingRule#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// One of MARKUP, DISCOUNT or TIERING that describes the behaviour of the pricing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#type BillingconductorPricingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces.
	//
	// Supported billing entities are AWS, AWS Marketplace, and AISPL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#billing_entity BillingconductorPricingRule#billing_entity}
	BillingEntity *string `field:"optional" json:"billingEntity" yaml:"billingEntity"`
	// Pricing rule description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#description BillingconductorPricingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Pricing rule modifier percentage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#modifier_percentage BillingconductorPricingRule#modifier_percentage}
	ModifierPercentage *float64 `field:"optional" json:"modifierPercentage" yaml:"modifierPercentage"`
	// The Operation which a SKU pricing rule is modifying.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#operation BillingconductorPricingRule#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// The service which a pricing rule is applied on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#service BillingconductorPricingRule#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#tags BillingconductorPricingRule#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The set of tiering configurations for the pricing rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#tiering BillingconductorPricingRule#tiering}
	Tiering *BillingconductorPricingRuleTiering `field:"optional" json:"tiering" yaml:"tiering"`
	// The UsageType which a SKU pricing rule is modifying.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_pricing_rule#usage_type BillingconductorPricingRule#usage_type}
	UsageType *string `field:"optional" json:"usageType" yaml:"usageType"`
}

