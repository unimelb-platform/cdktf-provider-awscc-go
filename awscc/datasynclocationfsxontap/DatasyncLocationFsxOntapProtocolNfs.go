package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolNfs struct {
	// The NFS mount options that DataSync can use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#mount_options DatasyncLocationFsxOntap#mount_options}
	MountOptions *DatasyncLocationFsxOntapProtocolNfsMountOptions `field:"required" json:"mountOptions" yaml:"mountOptions"`
}

