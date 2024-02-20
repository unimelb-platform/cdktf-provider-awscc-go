package memorydbparametergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemorydbParameterGroupConfig struct {
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
	// The name of the parameter group family that this parameter group is compatible with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#family MemorydbParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#parameter_group_name MemorydbParameterGroup#parameter_group_name}
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
	// A description of the parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#description MemorydbParameterGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An map of parameter names and values for the parameter update.
	//
	// You must supply at least one parameter name and value; subsequent arguments are optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#parameters MemorydbParameterGroup#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#tags MemorydbParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

