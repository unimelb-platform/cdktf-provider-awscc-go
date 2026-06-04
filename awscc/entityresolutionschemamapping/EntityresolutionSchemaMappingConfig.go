package entityresolutionschemamapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EntityresolutionSchemaMappingConfig struct {
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
	// The SchemaMapping attributes input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_schema_mapping#mapped_input_fields EntityresolutionSchemaMapping#mapped_input_fields}
	MappedInputFields interface{} `field:"required" json:"mappedInputFields" yaml:"mappedInputFields"`
	// The name of the SchemaMapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_schema_mapping#schema_name EntityresolutionSchemaMapping#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// The description of the SchemaMapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_schema_mapping#description EntityresolutionSchemaMapping#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_schema_mapping#tags EntityresolutionSchemaMapping#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

