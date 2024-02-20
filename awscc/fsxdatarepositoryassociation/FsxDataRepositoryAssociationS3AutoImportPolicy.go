package fsxdatarepositoryassociation


type FsxDataRepositoryAssociationS3AutoImportPolicy struct {
	// The ``AutoImportPolicy`` can have the following event values:   +   ``NEW`` - Amazon FSx automatically imports metadata of files added to the linked S3 bucket that do not currently exist in the FSx file system.
	//
	// +   ``CHANGED`` - Amazon FSx automatically updates file metadata and invalidates existing file content on the file system as files change in the data repository.
	//   +   ``DELETED`` - Amazon FSx automatically deletes files on the file system as corresponding files are deleted in the data repository.
	//
	//  You can define any combination of event types for your ``AutoImportPolicy``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#events FsxDataRepositoryAssociation#events}
	Events *[]*string `field:"required" json:"events" yaml:"events"`
}

