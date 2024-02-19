package cloudfrontcontinuousdeploymentpolicy


type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#enabled CloudfrontContinuousDeploymentPolicy#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#staging_distribution_dns_names CloudfrontContinuousDeploymentPolicy#staging_distribution_dns_names}.
	StagingDistributionDnsNames *[]*string `field:"required" json:"stagingDistributionDnsNames" yaml:"stagingDistributionDnsNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#single_header_policy_config CloudfrontContinuousDeploymentPolicy#single_header_policy_config}.
	SingleHeaderPolicyConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfig `field:"optional" json:"singleHeaderPolicyConfig" yaml:"singleHeaderPolicyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#single_weight_policy_config CloudfrontContinuousDeploymentPolicy#single_weight_policy_config}.
	SingleWeightPolicyConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfig `field:"optional" json:"singleWeightPolicyConfig" yaml:"singleWeightPolicyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#traffic_config CloudfrontContinuousDeploymentPolicy#traffic_config}.
	TrafficConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfig `field:"optional" json:"trafficConfig" yaml:"trafficConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#type CloudfrontContinuousDeploymentPolicy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

