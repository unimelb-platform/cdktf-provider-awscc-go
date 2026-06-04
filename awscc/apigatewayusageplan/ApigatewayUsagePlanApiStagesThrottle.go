package apigatewayusageplan


type ApigatewayUsagePlanApiStagesThrottle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_usage_plan#burst_limit ApigatewayUsagePlan#burst_limit}.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_usage_plan#rate_limit ApigatewayUsagePlan#rate_limit}.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

