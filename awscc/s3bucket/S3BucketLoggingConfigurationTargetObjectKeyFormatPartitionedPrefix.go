package s3bucket


type S3BucketLoggingConfigurationTargetObjectKeyFormatPartitionedPrefix struct {
	// Date Source for creating a partitioned prefix. This can be event time or delivery time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#partition_date_source S3Bucket#partition_date_source}
	PartitionDateSource *string `field:"optional" json:"partitionDateSource" yaml:"partitionDateSource"`
}

