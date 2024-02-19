package maciefindingsfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MacieFindingsFilterConfig struct {
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
	// Findings filter criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#finding_criteria MacieFindingsFilter#finding_criteria}
	FindingCriteria *MacieFindingsFilterFindingCriteria `field:"required" json:"findingCriteria" yaml:"findingCriteria"`
	// Findings filter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#name MacieFindingsFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Findings filter action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#action MacieFindingsFilter#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Findings filter description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#description MacieFindingsFilter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Findings filter position.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#position MacieFindingsFilter#position}
	Position *float64 `field:"optional" json:"position" yaml:"position"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#tags MacieFindingsFilter#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

