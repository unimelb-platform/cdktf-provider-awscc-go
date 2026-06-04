package fsxs3accesspointattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxS3AccessPointAttachmentConfig struct {
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
	// The Name of the S3AccessPointAttachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#name FsxS3AccessPointAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#open_zfs_configuration FsxS3AccessPointAttachment#open_zfs_configuration}.
	OpenZfsConfiguration *FsxS3AccessPointAttachmentOpenZfsConfiguration `field:"required" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#s3_access_point FsxS3AccessPointAttachment#s3_access_point}.
	S3AccessPoint *FsxS3AccessPointAttachmentS3AccessPoint `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
}

