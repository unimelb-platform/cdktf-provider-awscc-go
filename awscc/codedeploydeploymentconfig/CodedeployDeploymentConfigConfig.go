package codedeploydeploymentconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodedeployDeploymentConfigConfig struct {
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
	// The destination platform type for the deployment (Lambda, Server, or ECS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#compute_platform CodedeployDeploymentConfig#compute_platform}
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A name for the deployment configuration.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#deployment_config_name CodedeployDeploymentConfig#deployment_config_name}
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The minimum number of healthy instances that should be available at any time during the deployment.
	//
	// There are two parameters expected in the input: type and value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#minimum_healthy_hosts CodedeployDeploymentConfig#minimum_healthy_hosts}
	MinimumHealthyHosts *CodedeployDeploymentConfigMinimumHealthyHosts `field:"optional" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// The configuration that specifies how the deployment traffic is routed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#traffic_routing_config CodedeployDeploymentConfig#traffic_routing_config}
	TrafficRoutingConfig *CodedeployDeploymentConfigTrafficRoutingConfig `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
	// The zonal deployment config that specifies how the zonal deployment behaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#zonal_config CodedeployDeploymentConfig#zonal_config}
	ZonalConfig *CodedeployDeploymentConfigZonalConfig `field:"optional" json:"zonalConfig" yaml:"zonalConfig"`
}

