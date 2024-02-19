package datasynclocations3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationS3Config struct {
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
	// The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#s3_config DatasyncLocationS3#s3_config}
	S3Config *DatasyncLocationS3S3Config `field:"required" json:"s3Config" yaml:"s3Config"`
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#s3_bucket_arn DatasyncLocationS3#s3_bucket_arn}
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon S3 storage class you want to store your files in when this location is used as a task destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#s3_storage_class DatasyncLocationS3#s3_storage_class}
	S3StorageClass *string `field:"optional" json:"s3StorageClass" yaml:"s3StorageClass"`
	// A subdirectory in the Amazon S3 bucket.
	//
	// This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#subdirectory DatasyncLocationS3#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#tags DatasyncLocationS3#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

