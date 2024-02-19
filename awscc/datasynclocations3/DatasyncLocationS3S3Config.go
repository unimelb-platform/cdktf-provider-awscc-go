package datasynclocations3


type DatasyncLocationS3S3Config struct {
	// The ARN of the IAM role of the Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_s3#bucket_access_role_arn DatasyncLocationS3#bucket_access_role_arn}
	BucketAccessRoleArn *string `field:"required" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
}

