package s3storagelens


type S3StorageLensStorageLensConfigurationDataExportCloudwatchMetrics struct {
	// Specifies whether CloudWatch metrics are enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#is_enabled S3StorageLens#is_enabled}
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
}

