package codedeploydeploymentconfig


type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#canary_interval CodedeployDeploymentConfig#canary_interval}.
	CanaryInterval *float64 `field:"required" json:"canaryInterval" yaml:"canaryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#canary_percentage CodedeployDeploymentConfig#canary_percentage}.
	CanaryPercentage *float64 `field:"required" json:"canaryPercentage" yaml:"canaryPercentage"`
}

