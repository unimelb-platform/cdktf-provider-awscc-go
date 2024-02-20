package apigatewayv2apimapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2ApiMappingConfig struct {
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
	// The identifier of the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_api_mapping#api_id Apigatewayv2ApiMapping#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_api_mapping#domain_name Apigatewayv2ApiMapping#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The API stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_api_mapping#stage Apigatewayv2ApiMapping#stage}
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The API mapping key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_api_mapping#api_mapping_key Apigatewayv2ApiMapping#api_mapping_key}
	ApiMappingKey *string `field:"optional" json:"apiMappingKey" yaml:"apiMappingKey"`
}

