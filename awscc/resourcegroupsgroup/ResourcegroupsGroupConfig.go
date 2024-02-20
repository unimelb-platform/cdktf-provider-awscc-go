package resourcegroupsgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourcegroupsGroupConfig struct {
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
	// The name of the resource group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#name ResourcegroupsGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#configuration ResourcegroupsGroup#configuration}.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The description of the resource group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#description ResourcegroupsGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#resource_query ResourcegroupsGroup#resource_query}.
	ResourceQuery *ResourcegroupsGroupResourceQuery `field:"optional" json:"resourceQuery" yaml:"resourceQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#resources ResourcegroupsGroup#resources}.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourcegroups_group#tags ResourcegroupsGroup#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

