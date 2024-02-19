package datasynclocationfsxopenzfs


type DatasyncLocationFsxOpenZfsProtocol struct {
	// FSx OpenZFS file system NFS protocol information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs#nfs DatasyncLocationFsxOpenZfs#nfs}
	Nfs *DatasyncLocationFsxOpenZfsProtocolNfs `field:"optional" json:"nfs" yaml:"nfs"`
}

