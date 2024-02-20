package datasynclocationfsxopenzfs

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationFsxOpenZfsConfig struct {
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
	// Configuration settings for an NFS or SMB protocol, currently only support NFS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs#protocol DatasyncLocationFsxOpenZfs#protocol}
	Protocol *DatasyncLocationFsxOpenZfsProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs#security_group_arns DatasyncLocationFsxOpenZfs#security_group_arns}
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Name (ARN) for the FSx OpenZFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs#fsx_filesystem_arn DatasyncLocationFsxOpenZfs#fsx_filesystem_arn}
	FsxFilesystemArn *string `field:"optional" json:"fsxFilesystemArn" yaml:"fsxFilesystemArn"`
	// A subdirectory in the location's path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs#subdirectory DatasyncLocationFsxOpenZfs#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs#tags DatasyncLocationFsxOpenZfs#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

