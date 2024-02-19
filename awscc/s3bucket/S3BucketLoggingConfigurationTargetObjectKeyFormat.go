package s3bucket


type S3BucketLoggingConfigurationTargetObjectKeyFormat struct {
	// This format appends a time based prefix to the given log file prefix for delivering server access log file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#partitioned_prefix S3Bucket#partitioned_prefix}
	PartitionedPrefix *S3BucketLoggingConfigurationTargetObjectKeyFormatPartitionedPrefix `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// This format defaults the prefix to the given log file prefix for delivering server access log file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#simple_prefix S3Bucket#simple_prefix}
	SimplePrefix *string `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

