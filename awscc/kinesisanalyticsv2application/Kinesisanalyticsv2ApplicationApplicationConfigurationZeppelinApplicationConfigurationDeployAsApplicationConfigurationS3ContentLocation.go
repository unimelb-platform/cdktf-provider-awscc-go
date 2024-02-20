package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfigurationS3ContentLocation struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#bucket_arn Kinesisanalyticsv2Application#bucket_arn}
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The base path for the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#base_path Kinesisanalyticsv2Application#base_path}
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
}

