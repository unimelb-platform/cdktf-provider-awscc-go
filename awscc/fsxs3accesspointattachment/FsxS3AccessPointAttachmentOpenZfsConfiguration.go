package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#file_system_identity FsxS3AccessPointAttachment#file_system_identity}.
	FileSystemIdentity *FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentity `field:"required" json:"fileSystemIdentity" yaml:"fileSystemIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#volume_id FsxS3AccessPointAttachment#volume_id}.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}

