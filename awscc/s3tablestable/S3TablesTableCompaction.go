package s3tablestable


type S3TablesTableCompaction struct {
	// Indicates whether the Compaction maintenance action is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#status S3TablesTable#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The target file size for the table in MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table#target_file_size_mb S3TablesTable#target_file_size_mb}
	TargetFileSizeMb *float64 `field:"optional" json:"targetFileSizeMb" yaml:"targetFileSizeMb"`
}

