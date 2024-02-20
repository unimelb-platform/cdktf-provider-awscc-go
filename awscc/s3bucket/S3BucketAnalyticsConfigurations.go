package s3bucket


type S3BucketAnalyticsConfigurations struct {
	// The ID that identifies the analytics configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#storage_class_analysis S3Bucket#storage_class_analysis}
	StorageClassAnalysis *S3BucketAnalyticsConfigurationsStorageClassAnalysis `field:"required" json:"storageClassAnalysis" yaml:"storageClassAnalysis"`
	// The prefix that an object must have to be included in the analytics results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#tag_filters S3Bucket#tag_filters}.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

