package xrayresourcepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type XrayResourcePolicyConfig struct {
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
	// The resource policy document, which can be up to 5kb in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_resource_policy#policy_document XrayResourcePolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the resource policy. Must be unique within a specific AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_resource_policy#policy_name XrayResourcePolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// A flag to indicate whether to bypass the resource policy lockout safety check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_resource_policy#bypass_policy_lockout_check XrayResourcePolicy#bypass_policy_lockout_check}
	BypassPolicyLockoutCheck interface{} `field:"optional" json:"bypassPolicyLockoutCheck" yaml:"bypassPolicyLockoutCheck"`
}

