package s3bucket


type S3BucketAnalyticsConfigurationsStorageClassAnalysis struct {
	// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#data_export S3Bucket#data_export}
	DataExport *S3BucketAnalyticsConfigurationsStorageClassAnalysisDataExport `field:"optional" json:"dataExport" yaml:"dataExport"`
}

