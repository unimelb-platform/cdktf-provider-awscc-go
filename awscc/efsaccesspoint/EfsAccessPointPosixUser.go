package efsaccesspoint


type EfsAccessPointPosixUser struct {
	// The POSIX group ID used for all file system operations using this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#gid EfsAccessPoint#gid}
	Gid *string `field:"required" json:"gid" yaml:"gid"`
	// The POSIX user ID used for all file system operations using this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#uid EfsAccessPoint#uid}
	Uid *string `field:"required" json:"uid" yaml:"uid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#secondary_gids EfsAccessPoint#secondary_gids}
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

