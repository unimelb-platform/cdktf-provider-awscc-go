package apigatewaybasepathmapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayBasePathMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_base_path_mapping#domain_name ApigatewayBasePathMapping#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_base_path_mapping#base_path ApigatewayBasePathMapping#base_path}.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_base_path_mapping#rest_api_id ApigatewayBasePathMapping#rest_api_id}.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/apigateway_base_path_mapping#stage ApigatewayBasePathMapping#stage}.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

