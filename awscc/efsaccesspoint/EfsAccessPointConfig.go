package efsaccesspoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EfsAccessPointConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the EFS file system that the access point provides access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#file_system_id EfsAccessPoint#file_system_id}
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#access_point_tags EfsAccessPoint#access_point_tags}.
	AccessPointTags interface{} `field:"optional" json:"accessPointTags" yaml:"accessPointTags"`
	// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#client_token EfsAccessPoint#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The operating system user and group applied to all file system requests made using the access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#posix_user EfsAccessPoint#posix_user}
	PosixUser *EfsAccessPointPosixUser `field:"optional" json:"posixUser" yaml:"posixUser"`
	// Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point.
	//
	// The clients using the access point can only access the root directory and below. If the RootDirectory>Path specified does not exist, EFS creates it and applies the CreationInfo settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationInfo is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_access_point#root_directory EfsAccessPoint#root_directory}
	RootDirectory *EfsAccessPointRootDirectory `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
}

