package macieallowlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MacieAllowListConfig struct {
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
	// AllowList criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_allow_list#criteria MacieAllowList#criteria}
	Criteria *MacieAllowListCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Name of AllowList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_allow_list#name MacieAllowList#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of AllowList.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_allow_list#description MacieAllowList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_allow_list#tags MacieAllowList#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

