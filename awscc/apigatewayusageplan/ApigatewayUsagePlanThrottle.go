package apigatewayusageplan


type ApigatewayUsagePlanThrottle struct {
	// The API target request burst rate limit.
	//
	// This allows more requests through for a period of time than the target rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#burst_limit ApigatewayUsagePlan#burst_limit}
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API target request rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_usage_plan#rate_limit ApigatewayUsagePlan#rate_limit}
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

