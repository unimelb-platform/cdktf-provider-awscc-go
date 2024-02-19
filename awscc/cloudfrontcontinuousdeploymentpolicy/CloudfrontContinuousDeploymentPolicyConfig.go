package cloudfrontcontinuousdeploymentpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontContinuousDeploymentPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_continuous_deployment_policy#continuous_deployment_policy_config CloudfrontContinuousDeploymentPolicy#continuous_deployment_policy_config}.
	ContinuousDeploymentPolicyConfig *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfig `field:"required" json:"continuousDeploymentPolicyConfig" yaml:"continuousDeploymentPolicyConfig"`
}

