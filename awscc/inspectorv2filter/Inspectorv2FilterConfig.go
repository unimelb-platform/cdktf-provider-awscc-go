package inspectorv2filter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Inspectorv2FilterConfig struct {
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
	// Findings filter action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#filter_action Inspectorv2Filter#filter_action}
	FilterAction *string `field:"required" json:"filterAction" yaml:"filterAction"`
	// Findings filter criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#filter_criteria Inspectorv2Filter#filter_criteria}
	FilterCriteria *Inspectorv2FilterFilterCriteria `field:"required" json:"filterCriteria" yaml:"filterCriteria"`
	// Findings filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#name Inspectorv2Filter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Findings filter description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#description Inspectorv2Filter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

