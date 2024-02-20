package apigatewaydeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#rest_api_id ApigatewayDeployment#rest_api_id}
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The input configuration for a canary deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#deployment_canary_settings ApigatewayDeployment#deployment_canary_settings}
	DeploymentCanarySettings *ApigatewayDeploymentDeploymentCanarySettings `field:"optional" json:"deploymentCanarySettings" yaml:"deploymentCanarySettings"`
	// The description for the Deployment resource to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#description ApigatewayDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The description of the Stage resource for the Deployment resource to create.
	//
	// To specify a stage description, you must also provide a stage name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#stage_description ApigatewayDeployment#stage_description}
	StageDescription *ApigatewayDeploymentStageDescription `field:"optional" json:"stageDescription" yaml:"stageDescription"`
	// The name of the Stage resource for the Deployment resource to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#stage_name ApigatewayDeployment#stage_name}
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

