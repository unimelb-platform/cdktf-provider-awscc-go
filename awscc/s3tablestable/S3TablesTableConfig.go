package s3tablestable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3TablesTableConfig struct {
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
	// The namespace that the table belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#namespace S3TablesTable#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Format of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#open_table_format S3TablesTable#open_table_format}
	OpenTableFormat *string `field:"required" json:"openTableFormat" yaml:"openTableFormat"`
	// The Amazon Resource Name (ARN) of the specified table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#table_bucket_arn S3TablesTable#table_bucket_arn}
	TableBucketArn *string `field:"required" json:"tableBucketArn" yaml:"tableBucketArn"`
	// The name for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#table_name S3TablesTable#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Settings governing the Compaction maintenance action. Contains details about the compaction settings for an Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#compaction S3TablesTable#compaction}
	Compaction *S3TablesTableCompaction `field:"optional" json:"compaction" yaml:"compaction"`
	// Contains details about the metadata for an Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#iceberg_metadata S3TablesTable#iceberg_metadata}
	IcebergMetadata *S3TablesTableIcebergMetadata `field:"optional" json:"icebergMetadata" yaml:"icebergMetadata"`
	// Contains details about the snapshot management settings for an Iceberg table.
	//
	// A snapshot is expired when it exceeds MinSnapshotsToKeep and MaxSnapshotAgeHours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#snapshot_management S3TablesTable#snapshot_management}
	SnapshotManagement *S3TablesTableSnapshotManagement `field:"optional" json:"snapshotManagement" yaml:"snapshotManagement"`
	// Indicates that you don't want to specify a schema for the table.
	//
	// This property is mutually exclusive to 'IcebergMetadata', and its only possible value is 'Yes'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#without_metadata S3TablesTable#without_metadata}
	WithoutMetadata *string `field:"optional" json:"withoutMetadata" yaml:"withoutMetadata"`
}

