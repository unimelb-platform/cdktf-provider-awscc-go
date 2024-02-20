package resourceexplorer2view

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Resourceexplorer2ViewConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourceexplorer2_view#view_name Resourceexplorer2View#view_name}.
	ViewName *string `field:"required" json:"viewName" yaml:"viewName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourceexplorer2_view#filters Resourceexplorer2View#filters}.
	Filters *Resourceexplorer2ViewFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourceexplorer2_view#included_properties Resourceexplorer2View#included_properties}.
	IncludedProperties interface{} `field:"optional" json:"includedProperties" yaml:"includedProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourceexplorer2_view#scope Resourceexplorer2View#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resourceexplorer2_view#tags Resourceexplorer2View#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

