package apigatewaymethod


type ApigatewayMethodIntegrationIntegrationResponses struct {
	// Specifies the status code that is used to map the integration response to an existing MethodResponse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#status_code ApigatewayMethod#status_code}
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies how to handle response payload content type conversions.
	//
	// Supported values are ``CONVERT_TO_BINARY`` and ``CONVERT_TO_TEXT``, with the following behaviors:
	//  If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#content_handling ApigatewayMethod#content_handling}
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// A key-value map specifying response parameters that are passed to the method response from the back end.
	//
	// The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ``method.response.header.{name}``, where ``name`` is a valid and unique header name. The mapped non-static value must match the pattern of ``integration.response.header.{name}`` or ``integration.response.body.{JSON-expression}``, where ``name`` is a valid and unique response header name and ``JSON-expression`` is a valid JSON expression without the ``$`` prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#response_parameters ApigatewayMethod#response_parameters}
	ResponseParameters *map[string]*string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Specifies the templates used to transform the integration response body.
	//
	// Response templates are represented as a key/value map, with a content-type as the key and a template as the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#response_templates ApigatewayMethod#response_templates}
	ResponseTemplates *map[string]*string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Specifies the regular expression (regex) pattern used to choose an integration response based on the response from the back end.
	//
	// For example, if the success response returns nothing and the error response returns some string, you could use the ``.+`` regex to match error response. However, make sure that the error response does not contain any newline (``\n``) character in such cases. If the back end is an LAMlong function, the LAMlong function error header is matched. For all other HTTP and AWS back ends, the HTTP status code is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#selection_pattern ApigatewayMethod#selection_pattern}
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

