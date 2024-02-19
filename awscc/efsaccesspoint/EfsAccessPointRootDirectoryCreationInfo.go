package efsaccesspoint


type EfsAccessPointRootDirectoryCreationInfo struct {
	// Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#owner_gid EfsAccessPoint#owner_gid}
	OwnerGid *string `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#owner_uid EfsAccessPoint#owner_uid}
	OwnerUid *string `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#permissions EfsAccessPoint#permissions}
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

