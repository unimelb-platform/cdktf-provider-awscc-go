package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocation struct {
	// The Amazon Resource Name (ARN) for the S3 bucket containing the application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#bucket_arn Kinesisanalyticsv2Application#bucket_arn}
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The file key for the object containing the application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#file_key Kinesisanalyticsv2Application#file_key}
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// The version of the object containing the application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#object_version Kinesisanalyticsv2Application#object_version}
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

