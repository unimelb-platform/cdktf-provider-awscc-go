package apigatewaymethod


type ApigatewayMethodMethodResponses struct {
	// The method response's status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#status_code ApigatewayMethod#status_code}
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies the Model resources used for the response's content-type.
	//
	// Response models are represented as a key/value map, with a content-type as the key and a Model name as the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#response_models ApigatewayMethod#response_models}
	ResponseModels *map[string]*string `field:"optional" json:"responseModels" yaml:"responseModels"`
	// A key-value map specifying required or optional response parameters that API Gateway can send back to the caller.
	//
	// A key defines a method response header and the value specifies whether the associated method response header is required or not. The expression of the key must match the pattern ``method.response.header.{name}``, where ``name`` is a valid and unique header name. API Gateway passes certain integration response data to the method response headers specified here according to the mapping you prescribe in the API's IntegrationResponse. The integration response data that can be mapped include an integration response header expressed in ``integration.response.header.{name}``, a static value enclosed within a pair of single quotes (e.g., ``'application/json'``), or a JSON expression from the back-end response payload in the form of ``integration.response.body.{JSON-expression}``, where ``JSON-expression`` is a valid JSON expression without the ``$`` prefix.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#response_parameters ApigatewayMethod#response_parameters}
	ResponseParameters *map[string]interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

