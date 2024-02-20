package s3bucket


type S3BucketReplicationConfigurationRulesFilterAnd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tag_filters S3Bucket#tag_filters}.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

