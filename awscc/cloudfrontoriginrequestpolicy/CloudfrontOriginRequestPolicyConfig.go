package cloudfrontoriginrequestpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontOriginRequestPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#origin_request_policy_config CloudfrontOriginRequestPolicy#origin_request_policy_config}.
	OriginRequestPolicyConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfig `field:"required" json:"originRequestPolicyConfig" yaml:"originRequestPolicyConfig"`
}

