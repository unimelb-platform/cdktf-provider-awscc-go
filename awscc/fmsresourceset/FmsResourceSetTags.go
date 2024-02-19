package fmsresourceset


type FmsResourceSetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_resource_set#key FmsResourceSet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_resource_set#value FmsResourceSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

