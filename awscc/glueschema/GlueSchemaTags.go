package glueschema


type GlueSchemaTags struct {
	// A key to identify the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema#key GlueSchema#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Corresponding tag value for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema#value GlueSchema#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

