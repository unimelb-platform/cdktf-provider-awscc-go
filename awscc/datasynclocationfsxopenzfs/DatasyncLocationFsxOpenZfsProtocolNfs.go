package datasynclocationfsxopenzfs


type DatasyncLocationFsxOpenZfsProtocolNfs struct {
	// The NFS mount options that DataSync can use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_fsx_open_zfs#mount_options DatasyncLocationFsxOpenZfs#mount_options}
	MountOptions *DatasyncLocationFsxOpenZfsProtocolNfsMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

