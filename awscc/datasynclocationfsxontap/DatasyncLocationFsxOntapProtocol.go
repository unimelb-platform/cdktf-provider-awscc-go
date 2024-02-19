package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocol struct {
	// NFS protocol configuration for FSx ONTAP file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#nfs DatasyncLocationFsxOntap#nfs}
	Nfs *DatasyncLocationFsxOntapProtocolNfs `field:"optional" json:"nfs" yaml:"nfs"`
	// SMB protocol configuration for FSx ONTAP file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#smb DatasyncLocationFsxOntap#smb}
	Smb *DatasyncLocationFsxOntapProtocolSmb `field:"optional" json:"smb" yaml:"smb"`
}

