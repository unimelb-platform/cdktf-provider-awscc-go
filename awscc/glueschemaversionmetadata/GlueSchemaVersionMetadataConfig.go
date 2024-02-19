package glueschemaversionmetadata

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueSchemaVersionMetadataConfig struct {
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
	// Metadata key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema_version_metadata#key GlueSchemaVersionMetadata#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Represents the version ID associated with the schema version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema_version_metadata#schema_version_id GlueSchemaVersionMetadata#schema_version_id}
	SchemaVersionId *string `field:"required" json:"schemaVersionId" yaml:"schemaVersionId"`
	// Metadata value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/glue_schema_version_metadata#value GlueSchemaVersionMetadata#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

