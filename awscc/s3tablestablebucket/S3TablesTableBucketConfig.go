package s3tablestablebucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3TablesTableBucketConfig struct {
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
	// A name for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table_bucket#table_bucket_name S3TablesTableBucket#table_bucket_name}
	TableBucketName *string `field:"required" json:"tableBucketName" yaml:"tableBucketName"`
	// Specifies encryption settings for the table bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table_bucket#encryption_configuration S3TablesTableBucket#encryption_configuration}
	EncryptionConfiguration *S3TablesTableBucketEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Settings governing the Unreferenced File Removal maintenance action.
	//
	// Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3tables_table_bucket#unreferenced_file_removal S3TablesTableBucket#unreferenced_file_removal}
	UnreferencedFileRemoval *S3TablesTableBucketUnreferencedFileRemoval `field:"optional" json:"unreferencedFileRemoval" yaml:"unreferencedFileRemoval"`
}

