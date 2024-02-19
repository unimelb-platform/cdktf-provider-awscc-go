package eventschemasschema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventschemasSchemaConfig struct {
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
	// The source of the schema definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_schema#content EventschemasSchema#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The name of the schema registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_schema#registry_name EventschemasSchema#registry_name}
	RegistryName *string `field:"required" json:"registryName" yaml:"registryName"`
	// The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_schema#type EventschemasSchema#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_schema#description EventschemasSchema#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_schema#schema_name EventschemasSchema#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// Tags associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eventschemas_schema#tags EventschemasSchema#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

