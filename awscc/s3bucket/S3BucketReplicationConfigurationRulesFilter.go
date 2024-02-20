package s3bucket


type S3BucketReplicationConfigurationRulesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#and S3Bucket#and}.
	And *S3BucketReplicationConfigurationRulesFilterAnd `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Tags to use to identify a subset of objects for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tag_filter S3Bucket#tag_filter}
	TagFilter *S3BucketReplicationConfigurationRulesFilterTagFilter `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}

