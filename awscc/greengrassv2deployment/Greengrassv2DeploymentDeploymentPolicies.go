package greengrassv2deployment


type Greengrassv2DeploymentDeploymentPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#component_update_policy Greengrassv2Deployment#component_update_policy}.
	ComponentUpdatePolicy *Greengrassv2DeploymentDeploymentPoliciesComponentUpdatePolicy `field:"optional" json:"componentUpdatePolicy" yaml:"componentUpdatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#configuration_validation_policy Greengrassv2Deployment#configuration_validation_policy}.
	ConfigurationValidationPolicy *Greengrassv2DeploymentDeploymentPoliciesConfigurationValidationPolicy `field:"optional" json:"configurationValidationPolicy" yaml:"configurationValidationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#failure_handling_policy Greengrassv2Deployment#failure_handling_policy}.
	FailureHandlingPolicy *string `field:"optional" json:"failureHandlingPolicy" yaml:"failureHandlingPolicy"`
}

