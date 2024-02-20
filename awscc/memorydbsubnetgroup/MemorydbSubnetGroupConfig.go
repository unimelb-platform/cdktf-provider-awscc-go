package memorydbsubnetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemorydbSubnetGroupConfig struct {
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
	// The name of the subnet group.
	//
	// This value must be unique as it also serves as the subnet group identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_subnet_group#subnet_group_name MemorydbSubnetGroup#subnet_group_name}
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
	// A list of VPC subnet IDs for the subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_subnet_group#subnet_ids MemorydbSubnetGroup#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// An optional description of the subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_subnet_group#description MemorydbSubnetGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this subnet group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_subnet_group#tags MemorydbSubnetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

