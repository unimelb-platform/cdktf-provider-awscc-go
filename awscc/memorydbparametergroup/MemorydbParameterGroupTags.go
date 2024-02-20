package memorydbparametergroup


type MemorydbParameterGroupTags struct {
	// The key for the tag. May not be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#key MemorydbParameterGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag's value. May be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/memorydb_parameter_group#value MemorydbParameterGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

