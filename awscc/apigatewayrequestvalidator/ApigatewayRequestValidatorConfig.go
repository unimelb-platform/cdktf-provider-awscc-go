package apigatewayrequestvalidator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayRequestValidatorConfig struct {
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
	// The string identifier of the associated RestApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_request_validator#rest_api_id ApigatewayRequestValidator#rest_api_id}
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The name of this RequestValidator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_request_validator#name ApigatewayRequestValidator#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Boolean flag to indicate whether to validate a request body according to the configured Model schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_request_validator#validate_request_body ApigatewayRequestValidator#validate_request_body}
	ValidateRequestBody interface{} `field:"optional" json:"validateRequestBody" yaml:"validateRequestBody"`
	// A Boolean flag to indicate whether to validate request parameters (``true``) or not (``false``).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_request_validator#validate_request_parameters ApigatewayRequestValidator#validate_request_parameters}
	ValidateRequestParameters interface{} `field:"optional" json:"validateRequestParameters" yaml:"validateRequestParameters"`
}

