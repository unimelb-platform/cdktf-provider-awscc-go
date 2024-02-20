package apigatewayusageplan


type ApigatewayUsagePlanQuota struct {
	// The target maximum number of requests that can be made in a given time period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#limit ApigatewayUsagePlan#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The number of requests subtracted from the given limit in the initial time period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#offset ApigatewayUsagePlan#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#period ApigatewayUsagePlan#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
}

