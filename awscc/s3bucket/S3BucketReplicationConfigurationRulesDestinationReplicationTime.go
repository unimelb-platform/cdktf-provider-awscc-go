package s3bucket


type S3BucketReplicationConfigurationRulesDestinationReplicationTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#status S3Bucket#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#time S3Bucket#time}.
	Time *S3BucketReplicationConfigurationRulesDestinationReplicationTimeTime `field:"required" json:"time" yaml:"time"`
}

