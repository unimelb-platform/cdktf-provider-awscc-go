package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentityPosixUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#gid FsxS3AccessPointAttachment#gid}.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#uid FsxS3AccessPointAttachment#uid}.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#secondary_gids FsxS3AccessPointAttachment#secondary_gids}.
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

