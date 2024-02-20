package s3bucket


type S3BucketAnalyticsConfigurationsStorageClassAnalysisDataExport struct {
	// Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#destination S3Bucket#destination}
	Destination *S3BucketAnalyticsConfigurationsStorageClassAnalysisDataExportDestination `field:"required" json:"destination" yaml:"destination"`
	// The version of the output schema to use when exporting data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#output_schema_version S3Bucket#output_schema_version}
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

