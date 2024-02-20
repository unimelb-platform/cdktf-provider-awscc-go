package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevel struct {
	// Bucket-level metrics configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#bucket_level S3StorageLens#bucket_level}
	BucketLevel *S3StorageLensStorageLensConfigurationAccountLevelBucketLevel `field:"required" json:"bucketLevel" yaml:"bucketLevel"`
	// Enables activity metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#activity_metrics S3StorageLens#activity_metrics}
	ActivityMetrics *S3StorageLensStorageLensConfigurationAccountLevelActivityMetrics `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// Enables advanced cost optimization metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#advanced_cost_optimization_metrics S3StorageLens#advanced_cost_optimization_metrics}
	AdvancedCostOptimizationMetrics *S3StorageLensStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// Enables advanced data protection metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#advanced_data_protection_metrics S3StorageLens#advanced_data_protection_metrics}
	AdvancedDataProtectionMetrics *S3StorageLensStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// Enables detailed status codes metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#detailed_status_codes_metrics S3StorageLens#detailed_status_codes_metrics}
	DetailedStatusCodesMetrics *S3StorageLensStorageLensConfigurationAccountLevelDetailedStatusCodesMetrics `field:"optional" json:"detailedStatusCodesMetrics" yaml:"detailedStatusCodesMetrics"`
	// Specifies the details of Amazon S3 Storage Lens Group configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#storage_lens_group_level S3StorageLens#storage_lens_group_level}
	StorageLensGroupLevel *S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevel `field:"optional" json:"storageLensGroupLevel" yaml:"storageLensGroupLevel"`
}

