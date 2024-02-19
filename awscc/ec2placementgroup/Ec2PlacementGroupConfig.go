package ec2placementgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2PlacementGroupConfig struct {
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
	// The number of partitions. Valid only when **Strategy** is set to `partition`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_placement_group#partition_count Ec2PlacementGroup#partition_count}
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// The Spread Level of Placement Group is an enum where it accepts either host or rack when strategy is spread.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_placement_group#spread_level Ec2PlacementGroup#spread_level}
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
	// The placement strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_placement_group#strategy Ec2PlacementGroup#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_placement_group#tags Ec2PlacementGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

