package apigatewayv2deployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2DeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_deployment#api_id Apigatewayv2Deployment#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The description for the deployment resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_deployment#description Apigatewayv2Deployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of an existing stage to associate with the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigatewayv2_deployment#stage_name Apigatewayv2Deployment#stage_name}
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

