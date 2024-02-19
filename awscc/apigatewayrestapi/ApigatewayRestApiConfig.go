package apigatewayrestapi

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayRestApiConfig struct {
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
	// The source of the API key for metering requests according to a usage plan.
	//
	// Valid values are: ``HEADER`` to read the API key from the ``X-API-Key`` header of a request. ``AUTHORIZER`` to read the API key from the ``UsageIdentifierKey`` from a custom authorizer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#api_key_source_type ApigatewayRestApi#api_key_source_type}
	ApiKeySourceType *string `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#binary_media_types ApigatewayRestApi#binary_media_types}
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// An OpenAPI specification that defines a set of RESTful APIs in JSON format.
	//
	// For YAML templates, you can also provide the specification in YAML format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#body ApigatewayRestApi#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#body_s3_location ApigatewayRestApi#body_s3_location}
	BodyS3Location *ApigatewayRestApiBodyS3Location `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// The ID of the RestApi that you want to clone from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#clone_from ApigatewayRestApi#clone_from}
	CloneFrom *string `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// The description of the RestApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#description ApigatewayRestApi#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint.
	//
	// By default, clients can invoke your API with the default ``https://{api_id}.execute-api.{region}.amazonaws.com`` endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#disable_execute_api_endpoint ApigatewayRestApi#disable_execute_api_endpoint}
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ``Parameters`` property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#endpoint_configuration ApigatewayRestApi#endpoint_configuration}
	EndpointConfiguration *ApigatewayRestApiEndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// A query parameter to indicate whether to rollback the API update (``true``) or not (``false``) when a warning is encountered.
	//
	// The default value is ``false``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#fail_on_warnings ApigatewayRestApi#fail_on_warnings}
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API.
	//
	// When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#minimum_compression_size ApigatewayRestApi#minimum_compression_size}
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// This property applies only when you use OpenAPI to define your REST API.
	//
	// The ``Mode`` determines how API Gateway handles resource updates.
	//  Valid values are ``overwrite`` or ``merge``.
	//  For ``overwrite``, the new API definition replaces the existing one. The existing API identifier remains unchanged.
	//   For ``merge``, the new API definition is merged with the existing API.
	//  If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ``overwrite``. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API.
	//  Use the default mode to define top-level ``RestApi`` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#mode ApigatewayRestApi#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The name of the RestApi.
	//
	// A name is required if the REST API is not based on an OpenAPI specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#name ApigatewayRestApi#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Custom header parameters as part of the request.
	//
	// For example, to exclude DocumentationParts from an imported API, set ``ignore=documentation`` as a ``parameters`` value, as in the AWS CLI command of ``aws apigateway import-rest-api --parameters ignore=documentation --body 'file:///path/to/imported-api-body.json'``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#parameters ApigatewayRestApi#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for the ``RestApi`` resource.
	//
	// To set the ARN for the policy, use the ``!Join`` intrinsic function with ``""`` as delimiter and values of ``"execute-api:/"`` and ``"*"``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#policy ApigatewayRestApi#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The key-value map of strings.
	//
	// The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ``aws:``. The tag value can be up to 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_rest_api#tags ApigatewayRestApi#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

