package codedeploydeploymentconfig


type CodedeployDeploymentConfigTrafficRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#type CodedeployDeploymentConfig#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#time_based_canary CodedeployDeploymentConfig#time_based_canary}.
	TimeBasedCanary *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#time_based_linear CodedeployDeploymentConfig#time_based_linear}.
	TimeBasedLinear *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}

