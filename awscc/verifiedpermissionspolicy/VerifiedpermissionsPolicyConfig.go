package verifiedpermissionspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedpermissionsPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#definition VerifiedpermissionsPolicy#definition}.
	Definition *VerifiedpermissionsPolicyDefinition `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_policy#policy_store_id VerifiedpermissionsPolicy#policy_store_id}.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
}

