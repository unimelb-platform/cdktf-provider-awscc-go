package fsxdatarepositoryassociation


type FsxDataRepositoryAssociationS3AutoExportPolicy struct {
	// The ``AutoExportPolicy`` can have the following event values:   +   ``NEW`` - New files and directories are automatically exported to the data repository as they are added to the file system.
	//
	// +   ``CHANGED`` - Changes to files and directories on the file system are automatically exported to the data repository.
	//   +   ``DELETED`` - Files and directories are automatically deleted on the data repository when they are deleted on the file system.
	//
	//  You can define any combination of event types for your ``AutoExportPolicy``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#events FsxDataRepositoryAssociation#events}
	Events *[]*string `field:"required" json:"events" yaml:"events"`
}

