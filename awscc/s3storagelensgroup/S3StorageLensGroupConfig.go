package s3storagelensgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3StorageLensGroupConfig struct {
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
	// Sets the Storage Lens Group filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#filter S3StorageLensGroup#filter}
	Filter *S3StorageLensGroupFilter `field:"required" json:"filter" yaml:"filter"`
	// The name that identifies the Amazon S3 Storage Lens Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#name S3StorageLensGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of tags (key-value pairs) for this Amazon S3 Storage Lens Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens_group#tags S3StorageLensGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

