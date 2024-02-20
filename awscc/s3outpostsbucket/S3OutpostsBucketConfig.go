package s3outpostsbucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3OutpostsBucketConfig struct {
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
	// A name for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_bucket#bucket_name S3OutpostsBucket#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The id of the customer outpost on which the bucket resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_bucket#outpost_id S3OutpostsBucket#outpost_id}
	OutpostId *string `field:"required" json:"outpostId" yaml:"outpostId"`
	// Rules that define how Amazon S3Outposts manages objects during their lifetime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_bucket#lifecycle_configuration S3OutpostsBucket#lifecycle_configuration}
	LifecycleConfiguration *S3OutpostsBucketLifecycleConfiguration `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_bucket#tags S3OutpostsBucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

