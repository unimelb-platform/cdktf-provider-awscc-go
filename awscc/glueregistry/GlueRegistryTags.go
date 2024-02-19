package glueregistry


type GlueRegistryTags struct {
	// A key to identify the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_registry#key GlueRegistry#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Corresponding tag value for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_registry#value GlueRegistry#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

