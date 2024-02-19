package fsxdatarepositoryassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxDataRepositoryAssociationConfig struct {
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
	// The path to the Amazon S3 data repository that will be linked to the file system.
	//
	// The path can be an S3 bucket or prefix in the format ``s3://myBucket/myPrefix/``. This path specifies where in the S3 data repository files will be imported from or exported to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#data_repository_path FsxDataRepositoryAssociation#data_repository_path}
	DataRepositoryPath *string `field:"required" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// The ID of the file system on which the data repository association is configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#file_system_id FsxDataRepositoryAssociation#file_system_id}
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as ``/ns1/``) or subdirectory (such as ``/ns1/subdir/``) that will be mapped 1-1 with ``DataRepositoryPath``.
	//
	// The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path ``/ns1/``, then you cannot link another data repository with file system path ``/ns1/ns2``.
	//  This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.
	//   If you specify only a forward slash (``/``) as the file system path, you can link only one data repository to the file system. You can only specify "/" as the file system path for the first data repository associated with a file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#file_system_path FsxDataRepositoryAssociation#file_system_path}
	FileSystemPath *string `field:"required" json:"fileSystemPath" yaml:"fileSystemPath"`
	// A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created.
	//
	// The task runs if this flag is set to ``true``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#batch_import_meta_data_on_create FsxDataRepositoryAssociation#batch_import_meta_data_on_create}
	BatchImportMetaDataOnCreate interface{} `field:"optional" json:"batchImportMetaDataOnCreate" yaml:"batchImportMetaDataOnCreate"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.
	//  The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#imported_file_chunk_size FsxDataRepositoryAssociation#imported_file_chunk_size}
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association.
	//
	// The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#s3 FsxDataRepositoryAssociation#s3}
	S3 *FsxDataRepositoryAssociationS3 `field:"optional" json:"s3" yaml:"s3"`
	// An array of key-value pairs to apply to this resource.  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#tags FsxDataRepositoryAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

