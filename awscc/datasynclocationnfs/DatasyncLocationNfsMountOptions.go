package datasynclocationnfs


type DatasyncLocationNfsMountOptions struct {
	// The specific NFS version that you want DataSync to use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_nfs#version DatasyncLocationNfs#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

