package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolSmbMountOptions struct {
	// The specific SMB version that you want DataSync to use to mount your SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#version DatasyncLocationFsxOntap#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

