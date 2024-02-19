package apigatewayusageplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayUsagePlanConfig struct {
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
	// The associated API stages of a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#api_stages ApigatewayUsagePlan#api_stages}
	ApiStages interface{} `field:"optional" json:"apiStages" yaml:"apiStages"`
	// The description of a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#description ApigatewayUsagePlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The target maximum number of permitted requests per a given unit time interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#quota ApigatewayUsagePlan#quota}
	Quota *ApigatewayUsagePlanQuota `field:"optional" json:"quota" yaml:"quota"`
	// The collection of tags. Each tag element is associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#tags ApigatewayUsagePlan#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A map containing method level throttling information for API stage in a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#throttle ApigatewayUsagePlan#throttle}
	Throttle *ApigatewayUsagePlanThrottle `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#usage_plan_name ApigatewayUsagePlan#usage_plan_name}
	UsagePlanName *string `field:"optional" json:"usagePlanName" yaml:"usagePlanName"`
}

