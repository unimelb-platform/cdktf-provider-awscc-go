package apigatewayv2routeresponse

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2RouteResponseConfig struct {
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
	// The API identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_route_response#api_id Apigatewayv2RouteResponse#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The route ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_route_response#route_id Apigatewayv2RouteResponse#route_id}
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// The route response key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_route_response#route_response_key Apigatewayv2RouteResponse#route_response_key}
	RouteResponseKey *string `field:"required" json:"routeResponseKey" yaml:"routeResponseKey"`
	// The model selection expression for the route response. Supported only for WebSocket APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_route_response#model_selection_expression Apigatewayv2RouteResponse#model_selection_expression}
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// The response models for the route response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_route_response#response_models Apigatewayv2RouteResponse#response_models}
	ResponseModels *string `field:"optional" json:"responseModels" yaml:"responseModels"`
	// The route response parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigatewayv2_route_response#response_parameters Apigatewayv2RouteResponse#response_parameters}
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

