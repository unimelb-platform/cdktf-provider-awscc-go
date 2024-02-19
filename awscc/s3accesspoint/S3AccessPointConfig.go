package s3accesspoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3AccessPointConfig struct {
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
	// The name of the bucket that you want to associate this Access Point with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#bucket S3AccessPoint#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID associated with the S3 bucket associated with this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#bucket_account_id S3AccessPoint#bucket_account_id}
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The name you want to assign to this Access Point.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#name S3AccessPoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Access Point Policy you want to apply to this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#policy S3AccessPoint#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#public_access_block_configuration S3AccessPoint#public_access_block_configuration}
	PublicAccessBlockConfiguration *S3AccessPointPublicAccessBlockConfiguration `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// If you include this field, Amazon S3 restricts access to this Access Point to requests from the specified Virtual Private Cloud (VPC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_access_point#vpc_configuration S3AccessPoint#vpc_configuration}
	VpcConfiguration *S3AccessPointVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

