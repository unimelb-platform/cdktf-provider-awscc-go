package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#type CloudfrontContinuousDeploymentPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#single_header_config CloudfrontContinuousDeploymentPolicy#single_header_config}.
	SingleHeaderConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigSingleHeaderConfig `field:"optional" json:"singleHeaderConfig" yaml:"singleHeaderConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#single_weight_config CloudfrontContinuousDeploymentPolicy#single_weight_config}.
	SingleWeightConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigSingleWeightConfig `field:"optional" json:"singleWeightConfig" yaml:"singleWeightConfig"`
}

