package s3expressdirectorybucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ExpressDirectoryBucketConfig struct {
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
	// Specifies the number of Availability Zone that's used for redundancy for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3express_directory_bucket#data_redundancy S3ExpressDirectoryBucket#data_redundancy}
	DataRedundancy *string `field:"required" json:"dataRedundancy" yaml:"dataRedundancy"`
	// Specifies the AZ ID of the Availability Zone where the directory bucket will be created.
	//
	// An example AZ ID value is 'use1-az5'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3express_directory_bucket#location_name S3ExpressDirectoryBucket#location_name}
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
	// Specifies a name for the bucket.
	//
	// The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3express_directory_bucket#bucket_name S3ExpressDirectoryBucket#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
}

