package codedeploydeploymentconfig


type CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codedeploy_deployment_config#linear_interval CodedeployDeploymentConfig#linear_interval}.
	LinearInterval *float64 `field:"optional" json:"linearInterval" yaml:"linearInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/codedeploy_deployment_config#linear_percentage CodedeployDeploymentConfig#linear_percentage}.
	LinearPercentage *float64 `field:"optional" json:"linearPercentage" yaml:"linearPercentage"`
}

