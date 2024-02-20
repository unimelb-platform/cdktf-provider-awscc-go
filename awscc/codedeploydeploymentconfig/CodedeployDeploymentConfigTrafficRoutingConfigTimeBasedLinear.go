package codedeploydeploymentconfig


type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#linear_interval CodedeployDeploymentConfig#linear_interval}.
	LinearInterval *float64 `field:"required" json:"linearInterval" yaml:"linearInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codedeploy_deployment_config#linear_percentage CodedeployDeploymentConfig#linear_percentage}.
	LinearPercentage *float64 `field:"required" json:"linearPercentage" yaml:"linearPercentage"`
}

