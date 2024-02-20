package apigatewayusageplankey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayUsagePlanKeyConfig struct {
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
	// The Id of the UsagePlanKey resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan_key#key_id ApigatewayUsagePlanKey#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The type of a UsagePlanKey resource for a plan customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan_key#key_type ApigatewayUsagePlanKey#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan_key#usage_plan_id ApigatewayUsagePlanKey#usage_plan_id}
	UsagePlanId *string `field:"required" json:"usagePlanId" yaml:"usagePlanId"`
}

