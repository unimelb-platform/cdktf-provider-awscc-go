package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfigurationS3 struct {
	// Name of the Amazon S3 bucket where chat activity will be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_logging_configuration#bucket_name IvschatLoggingConfiguration#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
}

