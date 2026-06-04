package apigatewayv2routingrule


type Apigatewayv2RoutingRuleConditionsMatchHeadersAnyOf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_routing_rule#header Apigatewayv2RoutingRule#header}.
	Header *string `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_routing_rule#value_glob Apigatewayv2RoutingRule#value_glob}.
	ValueGlob *string `field:"optional" json:"valueGlob" yaml:"valueGlob"`
}

