package s3bucket


type S3BucketReplicationConfiguration struct {
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#role S3Bucket#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// A container for one or more replication rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

