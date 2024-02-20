package s3bucket


type S3BucketLifecycleConfigurationRulesTagFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#key S3Bucket#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#value S3Bucket#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

