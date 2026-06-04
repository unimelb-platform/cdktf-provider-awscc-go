package customerprofilessegmentdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesSegmentDefinitionConfig struct {
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
	// The display name of the segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_segment_definition#display_name CustomerprofilesSegmentDefinition#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The unique name of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_segment_definition#domain_name CustomerprofilesSegmentDefinition#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The unique name of the segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_segment_definition#segment_definition_name CustomerprofilesSegmentDefinition#segment_definition_name}
	SegmentDefinitionName *string `field:"required" json:"segmentDefinitionName" yaml:"segmentDefinitionName"`
	// An array that defines the set of segment criteria to evaluate when handling segment groups for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_segment_definition#segment_groups CustomerprofilesSegmentDefinition#segment_groups}
	SegmentGroups *CustomerprofilesSegmentDefinitionSegmentGroups `field:"required" json:"segmentGroups" yaml:"segmentGroups"`
	// The description of the segment definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_segment_definition#description CustomerprofilesSegmentDefinition#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_segment_definition#tags CustomerprofilesSegmentDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

