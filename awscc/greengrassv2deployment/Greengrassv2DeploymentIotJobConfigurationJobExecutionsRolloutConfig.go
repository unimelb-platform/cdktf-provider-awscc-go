package greengrassv2deployment


type Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#exponential_rate Greengrassv2Deployment#exponential_rate}.
	ExponentialRate *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigExponentialRate `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_deployment#maximum_per_minute Greengrassv2Deployment#maximum_per_minute}.
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

