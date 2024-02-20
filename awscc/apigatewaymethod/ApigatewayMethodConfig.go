package apigatewaymethod

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayMethodConfig struct {
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
	// The method's HTTP verb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#http_method ApigatewayMethod#http_method}
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The Resource identifier for the MethodResponse resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#resource_id ApigatewayMethod#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The string identifier of the associated RestApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#rest_api_id ApigatewayMethod#rest_api_id}
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// A boolean flag specifying whether a valid ApiKey is required to invoke this method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#api_key_required ApigatewayMethod#api_key_required}
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with a ``COGNITO_USER_POOLS`` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#authorization_scopes ApigatewayMethod#authorization_scopes}
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// The method's authorization type.
	//
	// This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference*.
	//   If you specify the ``AuthorizerId`` property, specify ``CUSTOM`` or ``COGNITO_USER_POOLS`` for this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#authorization_type ApigatewayMethod#authorization_type}
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The identifier of an authorizer to use on this method. The method's authorization type must be ``CUSTOM`` or ``COGNITO_USER_POOLS``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#authorizer_id ApigatewayMethod#authorizer_id}
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// Represents an ``HTTP``, ``HTTP_PROXY``, ``AWS``, ``AWS_PROXY``, or Mock integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#integration ApigatewayMethod#integration}
	Integration *ApigatewayMethodIntegration `field:"optional" json:"integration" yaml:"integration"`
	// Gets a method response associated with a given HTTP status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#method_responses ApigatewayMethod#method_responses}
	MethodResponses interface{} `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A human-friendly operation identifier for the method.
	//
	// For example, you can assign the ``operationName`` of ``ListPets`` for the ``GET /pets`` method in the ``PetStore`` example.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#operation_name ApigatewayMethod#operation_name}
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#request_models ApigatewayMethod#request_models}
	RequestModels *map[string]*string `field:"optional" json:"requestModels" yaml:"requestModels"`
	// A key-value map defining required or optional method request parameters that can be accepted by API Gateway.
	//
	// A key is a method request parameter name matching the pattern of ``method.request.{location}.{name}``, where ``location`` is ``querystring``, ``path``, or ``header`` and ``name`` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required (``true``) or optional (``false``). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#request_parameters ApigatewayMethod#request_parameters}
	RequestParameters *map[string]interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The identifier of a RequestValidator for request validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_method#request_validator_id ApigatewayMethod#request_validator_id}
	RequestValidatorId *string `field:"optional" json:"requestValidatorId" yaml:"requestValidatorId"`
}

