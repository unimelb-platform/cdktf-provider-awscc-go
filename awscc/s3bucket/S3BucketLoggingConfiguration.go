package s3bucket


type S3BucketLoggingConfiguration struct {
	// The name of an Amazon S3 bucket where Amazon S3 store server access log files.
	//
	// You can store log files in any bucket that you own. By default, logs are stored in the bucket where the LoggingConfiguration property is defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#destination_bucket_name S3Bucket#destination_bucket_name}
	DestinationBucketName *string `field:"optional" json:"destinationBucketName" yaml:"destinationBucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#log_file_prefix S3Bucket#log_file_prefix}.
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
	// Describes the key format for server access log file in the target bucket.
	//
	// You can choose between SimplePrefix and PartitionedPrefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#target_object_key_format S3Bucket#target_object_key_format}
	TargetObjectKeyFormat *S3BucketLoggingConfigurationTargetObjectKeyFormat `field:"optional" json:"targetObjectKeyFormat" yaml:"targetObjectKeyFormat"`
}

