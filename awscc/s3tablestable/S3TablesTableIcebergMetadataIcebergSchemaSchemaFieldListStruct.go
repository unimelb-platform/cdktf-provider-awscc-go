package s3tablestable


type S3TablesTableIcebergMetadataIcebergSchemaSchemaFieldListStruct struct {
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#name S3TablesTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A Boolean value that specifies whether values are required for each row in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#required S3TablesTable#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The field type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#type S3TablesTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

