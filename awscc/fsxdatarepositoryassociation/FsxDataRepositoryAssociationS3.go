package fsxdatarepositoryassociation


type FsxDataRepositoryAssociationS3 struct {
	// Describes a data repository association's automatic export policy.
	//
	// The ``AutoExportPolicy`` defines the types of updated objects on the file system that will be automatically exported to the data repository. As you create, modify, or delete files, Amazon FSx for Lustre automatically exports the defined changes asynchronously once your application finishes modifying the file.
	//  The ``AutoExportPolicy`` is only supported on Amazon FSx for Lustre file systems with a data repository association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#auto_export_policy FsxDataRepositoryAssociation#auto_export_policy}
	AutoExportPolicy *FsxDataRepositoryAssociationS3AutoExportPolicy `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// Describes the data repository association's automatic import policy.
	//
	// The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory listings up to date by importing changes to your Amazon FSx for Lustre file system as you modify objects in a linked S3 bucket.
	//  The ``AutoImportPolicy`` is only supported on Amazon FSx for Lustre file systems with a data repository association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#auto_import_policy FsxDataRepositoryAssociation#auto_import_policy}
	AutoImportPolicy *FsxDataRepositoryAssociationS3AutoImportPolicy `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

