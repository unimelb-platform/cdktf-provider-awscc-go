package memorydbacl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemorydbAclConfig struct {
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
	// The name of the acl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_acl#acl_name MemorydbAcl#acl_name}
	AclName *string `field:"required" json:"aclName" yaml:"aclName"`
	// An array of key-value pairs to apply to this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_acl#tags MemorydbAcl#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// List of users associated to this acl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_acl#user_names MemorydbAcl#user_names}
	UserNames *[]*string `field:"optional" json:"userNames" yaml:"userNames"`
}

