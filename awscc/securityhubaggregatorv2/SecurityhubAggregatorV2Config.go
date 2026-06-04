package securityhubaggregatorv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubAggregatorV2Config struct {
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
	// The list of included Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_aggregator_v2#linked_regions SecurityhubAggregatorV2#linked_regions}
	LinkedRegions *[]*string `field:"required" json:"linkedRegions" yaml:"linkedRegions"`
	// Indicates to link a list of included Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_aggregator_v2#region_linking_mode SecurityhubAggregatorV2#region_linking_mode}
	RegionLinkingMode *string `field:"required" json:"regionLinkingMode" yaml:"regionLinkingMode"`
	// A key-value pair to associate with the Security Hub V2 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_aggregator_v2#tags SecurityhubAggregatorV2#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

