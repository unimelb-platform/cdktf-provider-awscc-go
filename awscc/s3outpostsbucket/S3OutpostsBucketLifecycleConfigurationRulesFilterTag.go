package s3outpostsbucket


type S3OutpostsBucketLifecycleConfigurationRulesFilterTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3outposts_bucket#key S3OutpostsBucket#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3outposts_bucket#value S3OutpostsBucket#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

