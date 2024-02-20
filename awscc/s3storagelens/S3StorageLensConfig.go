package s3storagelens

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3StorageLensConfig struct {
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
	// Specifies the details of Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#storage_lens_configuration S3StorageLens#storage_lens_configuration}
	StorageLensConfiguration *S3StorageLensStorageLensConfiguration `field:"required" json:"storageLensConfiguration" yaml:"storageLensConfiguration"`
	// A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#tags S3StorageLens#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

