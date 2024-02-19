package datasynclocationsmb


type DatasyncLocationSmbMountOptions struct {
	// The specific SMB version that you want DataSync to use to mount your SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#version DatasyncLocationSmb#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

