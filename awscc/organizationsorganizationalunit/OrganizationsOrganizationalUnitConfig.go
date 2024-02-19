package organizationsorganizationalunit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationsOrganizationalUnitConfig struct {
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
	// The friendly name of this OU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_organizational_unit#name OrganizationsOrganizationalUnit#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier (ID) of the parent root or OU that you want to create the new OU in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_organizational_unit#parent_id OrganizationsOrganizationalUnit#parent_id}
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// A list of tags that you want to attach to the newly created OU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_organizational_unit#tags OrganizationsOrganizationalUnit#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

