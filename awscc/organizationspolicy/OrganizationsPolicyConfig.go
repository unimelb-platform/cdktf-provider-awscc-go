package organizationspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationsPolicyConfig struct {
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
	// The Policy text content.
	//
	// For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_policy#content OrganizationsPolicy#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Name of the Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_policy#name OrganizationsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_policy#type OrganizationsPolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Human readable description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_policy#description OrganizationsPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags that you want to attach to the newly created policy.
	//
	// For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_policy#tags OrganizationsPolicy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_policy#target_ids OrganizationsPolicy#target_ids}
	TargetIds *[]*string `field:"optional" json:"targetIds" yaml:"targetIds"`
}

