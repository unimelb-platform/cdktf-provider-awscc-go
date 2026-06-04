package s3tablestable


type S3TablesTableIcebergMetadata struct {
	// Contains details about the schema for an Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#iceberg_schema S3TablesTable#iceberg_schema}
	IcebergSchema *S3TablesTableIcebergMetadataIcebergSchema `field:"optional" json:"icebergSchema" yaml:"icebergSchema"`
}

