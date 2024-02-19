package apigatewayusageplan


type ApigatewayUsagePlanApiStages struct {
	// API Id of the associated API stage in a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#api_id ApigatewayUsagePlan#api_id}
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// API stage name of the associated API stage in a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#stage ApigatewayUsagePlan#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Map containing method level throttling information for API stage in a usage plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#throttle ApigatewayUsagePlan#throttle}
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

