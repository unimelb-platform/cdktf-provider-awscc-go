package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#posix_user FsxS3AccessPointAttachment#posix_user}.
	PosixUser *FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentityPosixUser `field:"required" json:"posixUser" yaml:"posixUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

