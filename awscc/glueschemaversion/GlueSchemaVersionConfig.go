package glueschemaversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueSchemaVersionConfig struct {
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
	// Identifier for the schema where the schema version will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema_version#schema GlueSchemaVersion#schema}
	Schema *GlueSchemaVersionSchema `field:"required" json:"schema" yaml:"schema"`
	// Complete definition of the schema in plain-text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema_version#schema_definition GlueSchemaVersion#schema_definition}
	SchemaDefinition *string `field:"required" json:"schemaDefinition" yaml:"schemaDefinition"`
}

