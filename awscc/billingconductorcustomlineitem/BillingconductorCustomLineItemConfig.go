package billingconductorcustomlineitem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingconductorCustomLineItemConfig struct {
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
	// Billing Group ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#billing_group_arn BillingconductorCustomLineItem#billing_group_arn}
	BillingGroupArn *string `field:"required" json:"billingGroupArn" yaml:"billingGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#name BillingconductorCustomLineItem#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account which this custom line item will be charged to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#account_id BillingconductorCustomLineItem#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#billing_period_range BillingconductorCustomLineItem#billing_period_range}.
	BillingPeriodRange *BillingconductorCustomLineItemBillingPeriodRange `field:"optional" json:"billingPeriodRange" yaml:"billingPeriodRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#custom_line_item_charge_details BillingconductorCustomLineItem#custom_line_item_charge_details}.
	CustomLineItemChargeDetails *BillingconductorCustomLineItemCustomLineItemChargeDetails `field:"optional" json:"customLineItemChargeDetails" yaml:"customLineItemChargeDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#description BillingconductorCustomLineItem#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/billingconductor_custom_line_item#tags BillingconductorCustomLineItem#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

