package apigatewaydocumentationpart

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayDocumentationPartConfig struct {
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
	// The location of the targeted API entity of the to-be-created documentation part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_documentation_part#location ApigatewayDocumentationPart#location}
	Location *ApigatewayDocumentationPartLocation `field:"required" json:"location" yaml:"location"`
	// The new documentation content map of the targeted API entity.
	//
	// Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_documentation_part#properties ApigatewayDocumentationPart#properties}
	Properties *string `field:"required" json:"properties" yaml:"properties"`
	// The string identifier of the associated RestApi.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_documentation_part#rest_api_id ApigatewayDocumentationPart#rest_api_id}
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

