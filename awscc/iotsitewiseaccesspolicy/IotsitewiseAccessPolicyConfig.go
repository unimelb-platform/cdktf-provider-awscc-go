package iotsitewiseaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseAccessPolicyConfig struct {
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
	// The identity for this access policy. Choose either a user or a group but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_access_policy#access_policy_identity IotsitewiseAccessPolicy#access_policy_identity}
	AccessPolicyIdentity *IotsitewiseAccessPolicyAccessPolicyIdentity `field:"required" json:"accessPolicyIdentity" yaml:"accessPolicyIdentity"`
	// The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_access_policy#access_policy_permission IotsitewiseAccessPolicy#access_policy_permission}
	AccessPolicyPermission *string `field:"required" json:"accessPolicyPermission" yaml:"accessPolicyPermission"`
	// The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_access_policy#access_policy_resource IotsitewiseAccessPolicy#access_policy_resource}
	AccessPolicyResource *IotsitewiseAccessPolicyAccessPolicyResource `field:"required" json:"accessPolicyResource" yaml:"accessPolicyResource"`
}

